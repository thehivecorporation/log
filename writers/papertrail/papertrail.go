package papertrail

import (
	"fmt"
	"net"
	"sync"

	"bytes"
	"log/syslog"
	"os"
	"strings"
	"time"

	"github.com/go-logfmt/logfmt"
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"
)

type writerImpl struct {
	telemetry.Common
	*Config

	mu   sync.Mutex
	conn net.Conn
}

func (w *writerImpl) WriteLog(payload *log.Payload) {
	ts := time.Now().Format(time.Stamp)

	var buf bytes.Buffer

	enc := logfmt.NewEncoder(&buf)
	enc.EncodeKeyval("level", payload.Level)
	enc.EncodeKeyval("message", strings.Join(payload.Messages, ", "))

	for k, v := range payload.Fields {
		enc.EncodeKeyval(k, v)
	}

	enc.EndRecord()

	msg := []byte(fmt.Sprintf("<%d>%s %s %s[%d]: %s\n", syslog.LOG_KERN, ts, w.Hostname, w.Tag, os.Getpid(), buf.String()))

	w.mu.Lock()
	_, err := w.conn.Write(msg)
	w.mu.Unlock()

	if err != nil {
		log.WithError(err).Error("Could not send data to papertrail")
	}
}

type Config struct {
	Host string
	Port int

	// Application settings
	Hostname string // Hostname value
	Tag      string // Tag value
}

// New handler.
func New(config *Config) log.Writer {
	conn, err := net.Dial("udp", fmt.Sprintf("%s.papertrailapp.com:%d", config.Host, config.Port))
	if err != nil {
		log.WithError(err).Fatal("Could not dial papertrail")
	}

	return &writerImpl{
		Config: config,
		conn:   conn,
	}
}
