package json

import (
	"encoding/json"
	"fmt"

	"io"

	"github.com/thehivecorporation/log"
)

type writer struct {
	w io.Writer
}

func (w *writer) WriteLog(p *log.Payload) {
	byt, _ := json.Marshal(p)

	fmt.Fprintln(w.w, string(byt))
}

func New(w io.Writer) log.Writer {
	return &writer{w}
}
