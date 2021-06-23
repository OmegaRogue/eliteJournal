package journal

type MissionData struct {
	MissionID        int    `json:"MissionID"`
	Name             string `json:"Name"`
	PassengerMission bool   `json:"PassengerMission"`
	Expires          int    `json:"Expires"`
}

type Missions struct {
	Event
	Active   []MissionData `json:"Active"`
	Failed   []MissionData `json:"Failed"`
	Complete []MissionData `json:"Complete"`
}
