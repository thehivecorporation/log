package multi

import (
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"
)

type writerImpl struct {
	telemetry.Common
	writers []log.Writer
}

func (w *writerImpl) WriteLog(payload *log.Payload) {
	for _, wr := range w.writers {
		wr.WriteLog(payload)
	}
}

func New(ws ...log.Writer) log.Writer {
	return &writerImpl{
		writers: ws,
	}
}
