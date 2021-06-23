package journal

// LoadGame LoadGame
type LoadGame struct {
	Event
	FID           string  `json:"FID"`
	Commander     string  `json:"Commander"`
	Horizons      bool    `json:"Horizons"`
	Ship          string  `json:"Ship"`
	ShipLocalised string  `json:"Ship_Localised"`
	ShipID        int     `json:"ShipID"`
	ShipName      string  `json:"ShipName"`
	ShipIdent     string  `json:"ShipIdent"`
	FuelLevel     float64 `json:"FuelLevel"`
	FuelCapacity  float64 `json:"FuelCapacity"`
	GameMode      string  `json:"GameMode"`
	Group         string  `json:"Group"`
	Credits       Cost    `json:"Credits"`
	Loan          Cost    `json:"Loan"`
}
