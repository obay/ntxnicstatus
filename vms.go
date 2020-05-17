package main

type vms struct {
	Metadata struct {
		GrandTotalEntities int `json:"grand_total_entities"`
		TotalEntities      int `json:"total_entities"`
		Count              int `json:"count"`
		StartIndex         int `json:"start_index"`
		EndIndex           int `json:"end_index"`
	} `json:"metadata"`
	Entities []struct {
		Affinity struct {
			Policy    string   `json:"policy"`
			HostUuids []string `json:"host_uuids"`
		} `json:"affinity,omitempty"`
		AllowLiveMigrate bool   `json:"allow_live_migrate"`
		GpusAssigned     bool   `json:"gpus_assigned"`
		HaPriority       int    `json:"ha_priority"`
		HostUUID         string `json:"host_uuid,omitempty"`
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
		VMLogicalTimestamp int    `json:"vm_logical_timestamp"`
		Description        string `json:"description,omitempty"`
	} `json:"entities"`
}
