package journal

// LoadGame LoadGame
type LoadGame struct {
	Event
	FID           string  `json:"FID"`
	Commander     string  `json:"Commander"`
	Horizons      bool    `json:"Horizons"`
	Ship          string  `json:"Ship,omitempty"`
	ShipLocalised string  `json:"Ship_Localised,omitempty"`
	ShipID        int     `json:"ShipID,omitempty"`
	ShipName      string  `json:"ShipName,omitempty"`
	ShipIdent     string  `json:"ShipIdent,omitempty"`
	FuelLevel     float64 `json:"FuelLevel,omitempty"`
	FuelCapacity  float64 `json:"FuelCapacity,omitempty"`
	GameMode      string  `json:"GameMode"`
	Group         string  `json:"Group"`
	Credits       Cost    `json:"Credits"`
	Loan          Cost    `json:"Loan"`
}
