package memory

import (
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"
)

type writerImpl struct {
	telemetry.Common
	payloads []*log.Payload
}

func (w *writerImpl) WriteLog(payload *log.Payload) {
	w.payloads = append(w.payloads, payload)
}

func New() log.Writer {
	return &writerImpl{
		payloads: make([]*log.Payload, 0, 20),
	}
}
