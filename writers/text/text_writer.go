package text

import (
	"github.com/thehivecorporation/log"
	"io"
)

type writer struct {
	log.TextWriter
}

func New(w io.Writer) log.Writer {
	return &writer{TextWriter: log.TextWriter{IOWriter: w}}
}
