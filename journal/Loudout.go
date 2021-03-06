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
