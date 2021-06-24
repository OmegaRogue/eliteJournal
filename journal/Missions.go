/*
 * Copyright (c) 2021 OmegaRogue.
 *  This file is part of eliteJournal.
 *
 *     eliteJournal is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or any later version.
 *
 *     eliteJournal is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with eliteJournal.  If not, see <https://www.gnu.org/licenses/>.
 */

package journal

// MissionData is the Data of a Mission in a Journal Event
type MissionData struct {
	MissionID        int    `json:"MissionID"`
	Name             string `json:"Name"`
	PassengerMission bool   `json:"PassengerMission"`
	Expires          int    `json:"Expires"`
}

// Missions Missions
type Missions struct {
	Event
	Active   []MissionData `json:"Active"`
	Failed   []MissionData `json:"Failed"`
	Complete []MissionData `json:"Complete"`
}
