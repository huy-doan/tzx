package outputdata

type AozoraConnectOutputData struct {
	IsConnected bool   `json:"is_connected"`
	AuthURL     string `json:"auth_url"`
}
