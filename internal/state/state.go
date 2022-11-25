package state

type Terraform struct {
  Version          int    `json:"version"`
  TerraformVersion string `json:"terraform_version"`
  Serial           int    `json:"serial"`
  Lineage          string `json:"lineage"`
  Outputs          struct {
    Ips struct {
      Value []string      `json:"value"`
      Type  []interface{} `json:"type"`
    } `json:"ips"`
  } `json:"outputs"`
  Resources []struct {
    Mode      string `json:"mode"`
    Type      string `json:"type"`
    Name      string `json:"name"`
    Provider  string `json:"provider"`
    Instances []struct {
      SchemaVersion int `json:"schema_version"`
      Attributes    struct {
        Description     string `json:"description,omitempty"`
        GuestProperties *struct {
          Hostname   string `json:"hostname"`
          PublicKeys string `json:"public-keys"`
          UserData   string `json:"user-data"`
        } `json:"guest_properties,omitempty"`
        Href  string `json:"href,omitempty"`
        Id    string `json:"id"`
        Lease []struct {
          RuntimeLeaseInSec int `json:"runtime_lease_in_sec"`
          StorageLeaseInSec int `json:"storage_lease_in_sec"`
        } `json:"lease,omitempty"`
        Metadata struct {
          VmwareCloudDirectorUiBadgesBlue    string `json:"vmware.cloud.director.ui.badges.blue"`
          VmwareCloudDirectorUiBadgesGreen   string `json:"vmware.cloud.director.ui.badges.green"`
          VmwareCloudDirectorUiBadgesMagenta string `json:"vmware.cloud.director.ui.badges.magenta"`
          VmwareCloudDirectorUiBadgesOrange  string `json:"vmware.cloud.director.ui.badges.orange"`
          VmwareCloudDirectorUiBadgesRed     string `json:"vmware.cloud.director.ui.badges.red"`
        } `json:"metadata,omitempty"`
        Name               string      `json:"name,omitempty"`
        Org                interface{} `json:"org"`
        PowerOn            bool        `json:"power_on,omitempty"`
        Status             int         `json:"status,omitempty"`
        StatusText         string      `json:"status_text,omitempty"`
        Vdc                interface{} `json:"vdc"`
        IsFenced           bool        `json:"is_fenced,omitempty"`
        OrgNetworkName     string      `json:"org_network_name,omitempty"`
        RetainIpMacEnabled bool        `json:"retain_ip_mac_enabled,omitempty"`
        VappName           string      `json:"vapp_name,omitempty"`
        AcceptAllEulas     bool        `json:"accept_all_eulas,omitempty"`
        BootImage          interface{} `json:"boot_image"`
        CatalogName        string      `json:"catalog_name,omitempty"`
        ComputerName       string      `json:"computer_name,omitempty"`
        CpuCores           int         `json:"cpu_cores,omitempty"`
        CpuHotAddEnabled   bool        `json:"cpu_hot_add_enabled,omitempty"`
        CpuLimit           interface{} `json:"cpu_limit"`
        CpuPriority        string      `json:"cpu_priority,omitempty"`
        CpuReservation     interface{} `json:"cpu_reservation"`
        CpuShares          interface{} `json:"cpu_shares"`
        Cpus               int         `json:"cpus,omitempty"`
        Customization      []struct {
          AdminPassword                  string `json:"admin_password"`
          AllowLocalAdminPassword        bool   `json:"allow_local_admin_password"`
          AutoGeneratePassword           bool   `json:"auto_generate_password"`
          ChangeSid                      bool   `json:"change_sid"`
          Enabled                        bool   `json:"enabled"`
          Force                          bool   `json:"force"`
          Initscript                     string `json:"initscript"`
          JoinDomain                     bool   `json:"join_domain"`
          JoinDomainAccountOu            string `json:"join_domain_account_ou"`
          JoinDomainName                 string `json:"join_domain_name"`
          JoinDomainPassword             string `json:"join_domain_password"`
          JoinDomainUser                 string `json:"join_domain_user"`
          JoinOrgDomain                  bool   `json:"join_org_domain"`
          MustChangePasswordOnFirstLogin bool   `json:"must_change_password_on_first_login"`
          NumberOfAutoLogons             int    `json:"number_of_auto_logons"`
        } `json:"customization,omitempty"`
        Disk                         []interface{} `json:"disk,omitempty"`
        ExposeHardwareVirtualization bool          `json:"expose_hardware_virtualization,omitempty"`
        HardwareVersion              string        `json:"hardware_version,omitempty"`
        InternalDisk                 []struct {
          BusNumber       int    `json:"bus_number"`
          BusType         string `json:"bus_type"`
          DiskId          string `json:"disk_id"`
          Iops            int    `json:"iops"`
          SizeInMb        int    `json:"size_in_mb"`
          StorageProfile  string `json:"storage_profile"`
          ThinProvisioned bool   `json:"thin_provisioned"`
          UnitNumber      int    `json:"unit_number"`
        } `json:"internal_disk,omitempty"`
        Memory              int         `json:"memory,omitempty"`
        MemoryHotAddEnabled bool        `json:"memory_hot_add_enabled,omitempty"`
        MemoryLimit         interface{} `json:"memory_limit"`
        MemoryPriority      string      `json:"memory_priority,omitempty"`
        MemoryReservation   interface{} `json:"memory_reservation"`
        MemoryShares        interface{} `json:"memory_shares"`
        Network             []struct {
          AdapterType      string `json:"adapter_type"`
          Connected        bool   `json:"connected"`
          Ip               string `json:"ip"`
          IpAllocationMode string `json:"ip_allocation_mode"`
          IsPrimary        bool   `json:"is_primary"`
          Mac              string `json:"mac"`
          Name             string `json:"name"`
          Type             string `json:"type"`
        } `json:"network,omitempty"`
        NetworkDhcpWaitSeconds interface{} `json:"network_dhcp_wait_seconds"`
        OsType                 string      `json:"os_type,omitempty"`
        OverrideTemplateDisk   []struct {
          BusNumber      int    `json:"bus_number"`
          BusType        string `json:"bus_type"`
          Iops           int    `json:"iops"`
          SizeInMb       int    `json:"size_in_mb"`
          StorageProfile string `json:"storage_profile"`
          UnitNumber     int    `json:"unit_number"`
        } `json:"override_template_disk,omitempty"`
        PreventUpdatePowerOff bool        `json:"prevent_update_power_off,omitempty"`
        SizingPolicyId        string      `json:"sizing_policy_id,omitempty"`
        StorageProfile        string      `json:"storage_profile,omitempty"`
        TemplateName          string      `json:"template_name,omitempty"`
        VmNameInTemplate      interface{} `json:"vm_name_in_template"`
        VmType                string      `json:"vm_type,omitempty"`
      } `json:"attributes"`
      SensitiveAttributes []interface{} `json:"sensitive_attributes"`
      Private             string        `json:"private"`
      Dependencies        []string      `json:"dependencies,omitempty"`
      IndexKey            int           `json:"index_key,omitempty"`
    } `json:"instances"`
  } `json:"resources"`
}
