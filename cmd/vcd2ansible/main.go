package main

import (
  "context"
  "encoding/json"
  "flag"
  "os"
  "text/template"
  
  "github.com/minio/minio-go/v7"
  "github.com/minio/minio-go/v7/pkg/credentials"
  
  "github.com/egeneralov/vcd2ansible/internal/state"
)

var (
  from     = flag.String("from", "fs", "fs/s3")
  endpoint = flag.String("endpoint", "", "s3 endpoint")
  id       = flag.String("id", "", "s3 access key")
  secret   = flag.String("secret", "", "s3 secret key")
  bucket   = flag.String("bucket", "", "s3 bucket name")
  object   = flag.String("object", "", "s3 object path")
  path     = flag.String("path", "terraform.tfstate", "to terraform state v4 json file with vcd")
  secure   = flag.Bool("secure", false, "http/s")
  vms      = map[string][]VM{}
)

type VM struct {
  ID   string `json:"id"`
  Name string `json:"name"`
  Host string `json:"host"`
}

func init() {
  flag.Parse()
}

func fromS3(endpoint, id, secret, bucket, object string, secure bool) (*state.Terraform, error) {
  s3Client, err := minio.New(endpoint, &minio.Options{
    Creds:  credentials.NewStaticV4(id, secret, ""),
    Secure: secure,
  })
  if err != nil {
    return nil, err
  }
  reader, err := s3Client.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
  if err != nil {
    return nil, err
  }
  defer reader.Close()
  var s state.Terraform
  if err := json.NewDecoder(reader).Decode(&s); err != nil {
    return nil, err
  }
  return &s, nil
}

func fromFile(path string) (*state.Terraform, error) {
  var s state.Terraform
  f, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  if err := json.NewDecoder(f).Decode(&s); err != nil {
    return nil, err
  }
  return &s, nil
}

func main() {
  var (
    s   *state.Terraform
    err error
  )
  
  switch *from {
  case "s3":
    s, err = fromS3(*endpoint, *id, *secret, *bucket, *object, *secure)
  case "fs":
    s, err = fromFile(*path)
  default:
    panic("invalid -from: " + *from + ", allowed values: [s3, fs]")
  }
  if err != nil {
    panic(err)
  }
  
  for _, resource := range s.Resources {
    if resource.Type != "vcd_vapp_vm" {
      continue
    }
    for _, instance := range resource.Instances {
      if !instance.Attributes.PowerOn {
        continue
      }
      vapp := instance.Attributes.VappName
      vms[vapp] = append(vms[vapp], VM{
        ID:   instance.Attributes.Id,
        Name: instance.Attributes.Name,
        Host: instance.Attributes.Network[0].Ip,
      })
    }
  }
  t, err := template.New("todos").Parse(tpl)
  if err != nil {
    panic(err)
  }
  if err := t.Execute(os.Stdout, vms); err != nil {
    panic(err)
  }
}

const tpl = `[vmware:children]
{{- range $group_name, $members := . }}
{{ $group_name }}
{{- end }}
{{ range $group_name, $members := . -}}
[{{ $group_name }}]
{{- range $members }}
{{ .Name }} ansible_host={{ .Host }} ip={{ .Host }} vmware_id={{ .ID }}
{{- end }}
{{- end }}
`
