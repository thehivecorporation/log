package custom

import "github.com/thehivecorporation/log"

type writer struct {
	w log.Writer
}

func (w *writer) WriteLog(p *log.Payload) {
	w.WriteLog(p)
}

func New(w log.Writer) log.Writer {
	return &writer{w}
}
