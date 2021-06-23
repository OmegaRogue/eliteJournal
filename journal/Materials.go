package journal

// Materials Materials
type Materials struct {
	Event
	Raw          []LocalizedValue `json:"Raw"`
	Manufactured []LocalizedValue `json:"Manufactured"`
	Encoded      []LocalizedValue `json:"Encoded"`
}
