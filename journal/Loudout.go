package journal

// Loadout Loadout
type Loadout struct {
	Event
	Ship          string  `json:"Ship"`
	ShipID        int     `json:"ShipID"`
	ShipName      string  `json:"ShipName"`
	ShipIdent     string  `json:"ShipIdent"`
	HullValue     Cost    `json:"HullValue"`
	ModulesValue  Cost    `json:"ModulesValue"`
	HullHealth    float64 `json:"HullHealth"`
	UnladenMass   float64 `json:"UnladenMass"`
	CargoCapacity int     `json:"CargoCapacity"`
	MaxJumpRange  float64 `json:"MaxJumpRange"`
	FuelCapacity  struct {
		Main    float64 `json:"Main"`
		Reserve float64 `json:"Reserve"`
	} `json:"FuelCapacity"`
	Rebuy   Cost `json:"Rebuy"`
	Modules []struct {
		Slot         string  `json:"Slot"`
		Item         string  `json:"Item"`
		On           bool    `json:"On"`
		Priority     int     `json:"Priority"`
		Health       float64 `json:"Health"`
		Value        Cost    `json:"Value,omitempty"`
		AmmoInClip   int     `json:"AmmoInClip,omitempty"`
		AmmoInHopper int     `json:"AmmoInHopper,omitempty"`
		Engineering  struct {
			BlueprintName      string  `json:"BlueprintName"`
			Level              int     `json:"Level"`
			Quality            float64 `json:"Quality"`
			EngineerID         int     `json:"EngineerID"`
			BlueprintID        int     `json:"blueprintID"`
			Engineer           string  `json:"Engineer"`
			ExperimentalEffect string  `json:"ExperimentalEffect,omitempty"`
			Modifiers          []struct {
				Label         string  `json:"Label"`
				Value         float64 `json:"Value,omitempty"`
				OriginalValue float64 `json:"OriginalValue"`
				LessIsGood    int     `json:"LessIsGood,omitempty"`
			} `json:"Modifiers"`
		} `json:"Engineering,omitempty"`
	} `json:"Modules"`
}
