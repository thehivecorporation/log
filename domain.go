package log

import "time"

// colors.
const (
	none   = 0
	red    = 31
	green  = 32
	yellow = 33
	blue   = 34
	gray   = 37
)

// Log levels.
const (
	InvalidLevel Level = iota - 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// Colors mapping.
var Colors = [...]int{
	LevelDebug: gray,
	LevelInfo:  blue,
	LevelWarn:  yellow,
	LevelError: red,
	LevelFatal: red,
}

var LevelNames = [...]string{
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warn",
	LevelError: "error",
	LevelFatal: "fatal",
}

var levelStrings = map[string]Level{
	"debug":   LevelDebug,
	"info":    LevelInfo,
	"warn":    LevelWarn,
	"warning": LevelWarn,
	"error":   LevelError,
	"fatal":   LevelFatal,
}

type Tags []string
type Fields map[string]interface{}
type Level int

type Logger interface {
	Debug(msg string) Telemetry
	Info(msg string) Telemetry
	Warn(msg string) Telemetry
	Error(msg string) Telemetry
	Fatal(msg string) Telemetry

	Debugf(msg string, v ...interface{}) Telemetry
	Infof(msg string, v ...interface{}) Telemetry
	Warnf(msg string, v ...interface{}) Telemetry
	Errorf(msg string, v ...interface{}) Telemetry
	Fatalf(msg string, v ...interface{}) Telemetry

	WithField(s string, v interface{}) Logger
	WithFields(Fields) Logger
	WithError(error) Logger

	WithTags(s ...string) Telemetry

	Clone() Logger
}

type Payload struct {
	Level     Level      `json:"level,omitempty"`
	Messages  []string   `json:"messages,omitempty"`
	Fields    Fields     `json:"fields,omitempty"`
	Timestamp *time.Time `json:"ts,omitempty"`
	Tags      Tags       `json:"tags,omitempty"`
	errors    []string   `json:"errors,omitempty"`
}

type Writer interface {
	WriteLog(payload *Payload)
}

type Telemetry interface {
	WithTags(s ...string) Telemetry

	Inc(name string, value float64, extra ...interface{}) Logger
	Gauge(string, float64, ...interface{}) Logger
	Histogram(name string, value float64, extra ...interface{}) Logger

	Clone() Telemetry
	SetLogger(l Logger)
}
