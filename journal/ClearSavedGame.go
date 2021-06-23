package journal

// ClearSavedGame ClearSavedGame
type ClearSavedGame struct {
	Event
	Name string `json:"Name"`
	FID  string `json:"FID"`
}
