package journal

// Commander Commander
type Commander struct {
	Event
	FID  string `json:"FID"`
	Name string `json:"Name"`
}
