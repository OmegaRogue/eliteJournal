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
