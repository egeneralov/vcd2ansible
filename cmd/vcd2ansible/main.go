package main

import (
  "encoding/json"
  "flag"
  "os"
  "text/template"
  
  "github.com/egeneralov/vcd2ansible/internal/state"
)

var (
  path = flag.String("path", "terraform.tfstate", "path to terraform state v4 json file with vcd")
  s    state.Terraform
  vms  = map[string][]VM{}
)

type VM struct {
  ID   string `json:"id"`
  Name string `json:"name"`
  Host string `json:"host"`
}

func init() {
  flag.Parse()
}

func main() {
  f, err := os.Open(*path)
  if err != nil {
    panic(err)
  }
  if err := json.NewDecoder(f).Decode(&s); err != nil {
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
