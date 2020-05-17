package main

type vmDetails struct {
	AllowLiveMigrate bool   `json:"allow_live_migrate"`
	GpusAssigned     bool   `json:"gpus_assigned"`
	Description      string `json:"description"`
	HaPriority       int    `json:"ha_priority"`
	HostUUID         string `json:"host_uuid"`
	MemoryMb         int    `json:"memory_mb"`
	Name             string `json:"name"`
	NumCoresPerVcpu  int    `json:"num_cores_per_vcpu"`
	NumVcpus         int    `json:"num_vcpus"`
	PowerState       string `json:"power_state"`
	Timezone         string `json:"timezone"`
	UUID             string `json:"uuid"`
	VMFeatures       struct {
		VGACONSOLE bool `json:"VGA_CONSOLE"`
		AGENTVM    bool `json:"AGENT_VM"`
	} `json:"vm_features"`
	VMLogicalTimestamp int `json:"vm_logical_timestamp"`
}
