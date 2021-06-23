//go:generate go-enum -f=$GOFILE --marshal --noprefix --sql --sqlnullint

package elite

// Allegiance is an enumeration of Elite Dangerous Allegiances
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
