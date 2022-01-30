package utilities

import (
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func Initialize() {
	Cron = cron.New()
}
