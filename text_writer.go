package log

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type TextWriter struct {
	IOWriter io.Writer
}

func newTextWriter(w io.Writer) Writer {
	return &TextWriter{IOWriter: w}
}

func (w *TextWriter) WriteLog(p *Payload) {
	ts := time.Since(*p.Timestamp) / time.Second

	for _, msg := range p.Messages {
		fmt.Fprintf(w.IOWriter, "\033[%dm%6s\033[0m[%04d] %-25s", Colors[p.Level], strings.ToUpper(LevelNames[p.Level]), ts, msg)
	}

	for k, value := range p.Fields {
		fmt.Fprintf(w.IOWriter, " \033[%dm%s\033[0m=%v", Colors[p.Level], k, value)
	}

	if len(p.Tags) > 0 {
		fmt.Fprintf(w.IOWriter, " \033[%dm%s\033[0m", Colors[p.Level], strings.Join(p.Tags, "|"))
	}

	fmt.Fprintln(w.IOWriter)

	for _, msg := range p.errors {
		fmt.Fprintf(w.IOWriter, "\033[%dm%6s\033[0m[%04d]    %-25s\n", Colors[p.Level], strings.ToUpper(LevelNames[p.Level]), ts, msg)
	}
}
