package journal

type NewCommander struct {
	Event
	Name    string `json:"Name"`
	FID     string `json:"FID"`
	Package string `json:"Package"`
}
