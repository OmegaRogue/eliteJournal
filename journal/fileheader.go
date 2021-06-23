package journal

// FileHeader fileheader
type FileHeader struct {
	Event
	Part        int    `json:"part"`
	Language    string `json:"language"`
	Odyssey     bool   `json:"Odyssey,omitempty"`
	GameVersion string `json:"gameversion"`
	Build       string `json:"build"`
}
