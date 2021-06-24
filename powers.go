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

//go:generate go-enum -f=$GOFILE --marshal --noprefix --sqlnullint

package elite

// Power is an enumeration of Elite Dangerous Powers
/*
ENUM(
Uncontrolled
Aisling Duval
Archon Delaine
Arissa Lavigny-Duval
Denton Patreus
Edmund Mahon
Felicia Winters
Li Yong-Rui
Pranav Antal
Yuri Grom
Zachary Hudson
Zemina Torval
)
*/
type Power int
