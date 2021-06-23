//go:generate go-enum -f=$GOFILE --marshal --noprefix --sqlnullint

package elite

// Allegiance is an enumeration of Elite Dangerous Allegiances with Major Factions
/*
ENUM(
Independent
Alliance
Empire
Federation
Pirate
PilotsFederation
Thargoids
Guardians
)
*/
type Allegiance int
