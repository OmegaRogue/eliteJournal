//go:generate go-enum -f=$GOFILE --marshal --noprefix --sql --sqlnullint
package elite

// Rank is an enumeration of Elite Dangerous Ranks
/*
ENUM(
None=-1
Rookie
Agent
Officer
SeniorOfficer
Leader
)
*/
type Rank int
