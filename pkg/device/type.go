package device

type Device struct {
	User    string `json:"user"`
	OS      string `json:"os"`
	Product string `json:"product"`
	Type    string `json:"type"`
}
