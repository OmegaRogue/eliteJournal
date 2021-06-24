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

package elite

import (
	"github.com/OmegaRogue/eliteJournal/journal"
)

// SLEF Ship data for inara and various other tools to commonly support
type SLEF struct {
	Header struct {
		AppName    string `json:"appName"`
		AppVersion int    `json:"appVersion"`
		AppURL     string `json:"appURL"`
	} `json:"header"`
	Data journal.Loadout `json:"data"`
}
