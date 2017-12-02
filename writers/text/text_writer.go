package text

import (
	"github.com/thehivecorporation/log"
	"io"
	"os"
)

type writer struct {
	log.TextWriter
}

func New(w ...io.Writer) log.Writer {
	switch len(w) {
	case 1:
		return &writer{TextWriter: log.TextWriter{IOWriter: w[0], ErrorIOWriter: w[0]}}
	case 2:
		return &writer{TextWriter: log.TextWriter{IOWriter: w[0], ErrorIOWriter: w[1]}}
	default:
		return &writer{TextWriter: log.TextWriter{IOWriter: os.Stdout, ErrorIOWriter: os.Stderr}}
	}
}
