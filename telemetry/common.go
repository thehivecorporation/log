package telemetry

import (
	"github.com/thehivecorporation/log"
	"time"
)

type Common struct {
	Logger log.Logger
	Tags   []string
	Start  time.Time
}