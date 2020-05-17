package main

type vmNICs struct {
	Metadata struct {
		GrandTotalEntities int `json:"grand_total_entities"`
		TotalEntities      int `json:"total_entities"`
	} `json:"metadata"`
	Entities []struct {
		MacAddress  string `json:"mac_address"`
		NetworkUUID string `json:"network_uuid"`
		Model       string `json:"model"`
		IPAddress   string `json:"ip_address"`
		IsConnected bool   `json:"is_connected"`
	} `json:"entities"`
}
