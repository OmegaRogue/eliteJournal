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
