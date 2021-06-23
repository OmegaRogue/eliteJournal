//go:generate go-enum -f=$GOFILE --marshal --noprefix --sql --sqlnullint

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
