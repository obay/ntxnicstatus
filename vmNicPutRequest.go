package main

type vmNicPutRequest struct {
	NicID   string `json:"nic_id"`
	NicSpec struct {
		// AdapterType        string `json:"adapter_type"`
		IsConnected bool `json:"is_connected"`
		// MacAddress         string `json:"mac_address"`
		// Model              string `json:"model"`
		NetworkUUID string `json:"network_uuid"`
		// RequestIP          bool   `json:"request_ip"`
		// RequestedIPAddress string `json:"requested_ip_address"`
	} `json:"nic_spec"`
	// UUID               string `json:"uuid"`
	// VMLogicalTimestamp int    `json:"vm_logical_timestamp"`
}
