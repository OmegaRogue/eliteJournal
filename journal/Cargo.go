package journal

// Cargo Cargo
type Cargo struct {
	Event
	Vessel    string      `json:"Vessel"`
	Count     int         `json:"Count"`
	Inventory []CargoItem `json:"Inventory"`
}
type CargoItem struct {
	MissionID int `json:"MissionID,omitempty"`
	LocalizedValue
	Stolen int `json:"Stolen"`
}
