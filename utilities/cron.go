package utilities

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func Initialize() {
	Cron = cron.New()
}

func AllScheduledEntries() {
	entries := Cron.Entries()

	for _, entry := range entries {
		fmt.Println(entry.ID)
	}
}
