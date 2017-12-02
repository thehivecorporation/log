package json

import (
	"encoding/json"
	"fmt"

	"io"

	"github.com/thehivecorporation/log"
	"os"
)

type writer struct {
	w io.Writer
	e io.Writer
}

func (w *writer) WriteLog(p *log.Payload) {
	byt, _ := json.Marshal(p)

	if len(p.Errors) > 0 {
		fmt.Fprintln(w.e, string(byt))
	} else {
		fmt.Fprintln(w.w, string(byt))
	}
}

func New(w ...io.Writer) log.Writer {
	switch len(w) {
	case 1:
		return &writer{w: w[0], e: w[0]}
	case 2:
		return &writer{w: w[0], e: w[1]}
	default:
		return &writer{w: os.Stdout, e: os.Stderr}
	}
}
