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

// Cargo Cargo
type Cargo struct {
	Event
	Vessel    string      `json:"Vessel"`
	Count     int         `json:"Count"`
	Inventory []CargoItem `json:"Inventory"`
}

// CargoItem represents a type of Commodity in Cargo
type CargoItem struct {
	MissionID int `json:"MissionID,omitempty"`
	LocalizedValue
	Stolen int `json:"Stolen"`
}
