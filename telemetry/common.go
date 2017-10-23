package telemetry

import (
	"github.com/thehivecorporation/log"
	"time"
)

type Common struct {
	Logger log.Logger
	Tags   log.Tags
	Start  time.Time
}
