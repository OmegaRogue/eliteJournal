package elite

import (
	"github.com/eliteJournal/journal"
)

type SLEF struct {
	Header struct {
		AppName    string `json:"appName"`
		AppVersion int    `json:"appVersion"`
		AppURL     string `json:"appURL"`
	} `json:"header"`
	Data journal.Loadout `json:"data"`
}
