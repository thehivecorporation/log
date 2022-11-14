package log

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
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
		return &TextWriter{IOWriter: color.Output, ErrorIOWriter: color.Error}
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

	writeMessage(writer, p)
	writeFields(writer, p)
	writeTags(writer, p)

	fmt.Fprintln(writer)

	writeErrors(writer, p)
}

func writeErrors(w io.Writer, p *Payload) {
	attr := Colors[p.Level]
	level := color.New(attr)

	for _, msg := range p.Errors {
		level.Fprintln(w, "%6s\033[0m[%04d]    %-25s", strings.ToUpper(LevelNames[p.Level]), p.ElapsedSinceStart, msg)
	}
}

func writeTags(w io.Writer, p *Payload) {
	attr := Colors[p.Level]
	level := color.New(attr)

	if len(p.Tags) > 0 {
		for k, v := range p.Tags {
			level.Fprintf(w, " %s|%s\033[0m", k, v)
		}
	}
}

func writeFields(w io.Writer, p *Payload) {
	attr := Colors[p.Level]
	level := color.New(attr)
	for k, v := range p.Fields {
		level.Fprintf(w, " %s\033[0m=%v", k, v)
	}
}

func writeMessage(w io.Writer, p *Payload) {
	attr := Colors[p.Level]
	level := color.New(attr)
	for _, msg := range p.Messages {
		level.Fprintf(w, "%6s", strings.ToUpper(LevelNames[p.Level]))
		fmt.Fprintf(w, "[%04d](%-25s) %-25s", p.ElapsedSinceStart, p.Timestamp.Format(time.RFC3339Nano), msg)
	}
}
