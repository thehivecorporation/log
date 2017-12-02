package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type TextWriter struct {
	sync.Mutex
	IOWriter      io.Writer
	ErrorIOWriter io.Writer
}

func newTextWriter(w ...io.Writer) Writer {
	switch len(w) {
	case 1:
		return &TextWriter{IOWriter: w[0], ErrorIOWriter: w[0]}
	case 2:
		return &TextWriter{IOWriter: w[0], ErrorIOWriter: w[1]}
	default:
		return &TextWriter{IOWriter: os.Stdout, ErrorIOWriter: os.Stderr}
	}

}

func (w *TextWriter) WriteLog(p *Payload) {
	var writer io.Writer
	if len(p.Errors) > 0 {
		writer = w.ErrorIOWriter
	} else {
		writer = w.IOWriter
	}

	w.Lock()
	defer w.Unlock()

	ts := time.Since(*p.Timestamp) / time.Second

	for _, msg := range p.Messages {
		fmt.Fprintf(writer, "\033[%dm%6s\033[0m[%04d] %-25s", Colors[p.Level], strings.ToUpper(LevelNames[p.Level]), ts, msg)
	}

	for k, value := range p.Fields {
		fmt.Fprintf(writer, " \033[%dm%s\033[0m=%v", Colors[p.Level], k, value)
	}

	if len(p.Tags) > 0 {
		for k, v := range p.Tags {
			fmt.Fprintf(writer, " \033[%dm%s|%s\033[0m", Colors[p.Level], k, v)
		}
	}

	fmt.Fprintln(writer)

	for _, msg := range p.Errors {
		fmt.Fprintf(writer, "\033[%dm%6s\033[0m[%04d]    %-25s\n", Colors[p.Level], strings.ToUpper(LevelNames[p.Level]), ts, msg)
	}
}
