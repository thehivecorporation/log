package log

import (
	"time"

	"github.com/fatih/color"
)

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
var Colors = [...]color.Attribute{
	LevelDebug: color.FgWhite,
	LevelInfo:  color.FgBlue,
	LevelWarn:  color.FgYellow,
	LevelError: color.FgRed,
	LevelFatal: color.FgRed,
}

var LevelNames = [...]string{
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warn",
	LevelError: "error",
	LevelFatal: "fatal",
}

var LevelStrings = map[string]Level{
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
}

func LevelFromString(s string) Level {
	return LevelStrings[s]
}

type Tags map[string]string
type Fields map[string]interface{}
type Level int

type Logger interface {
	Debug(msg interface{}, more ...interface{}) Telemetry
	Info(msg interface{}, more ...interface{}) Telemetry
	Warn(msg interface{}, more ...interface{}) Telemetry
	Error(msg interface{}, more ...interface{}) Telemetry
	Fatal(msg interface{}, more ...interface{}) Telemetry

	Debugf(msg string, v ...interface{}) Telemetry
	Infof(msg string, v ...interface{}) Telemetry
	Warnf(msg string, v ...interface{}) Telemetry
	Errorf(msg string, v ...interface{}) Telemetry
	Fatalf(msg string, v ...interface{}) Telemetry

	WithField(s string, v interface{}) Logger
	WithFields(Fields) Logger
	WithError(...error) Logger
	WithErrors(...error) Logger

	WithTags(t Tags) Telemetry
	WithTag(string, string) Telemetry

	Clone() Logger
}

type Payload struct {
	Level             Level         `json:"level,omitempty"`
	Messages          []interface{} `json:"messages,omitempty"`
	Fields            Fields        `json:"fields,omitempty"`
	Timestamp         time.Time     `json:"ts,omitempty"`
	Tags              Tags          `json:"tags,omitempty"`
	Errors            []string      `json:"Errors,omitempty"`
	ElapsedSinceStart time.Duration `json:"ElapsedSinceStart,omitempty"`
}

type Writer interface {
	WriteLog(payload *Payload)
}

type Telemetry interface {
	WithTags(t Tags) Telemetry
	WithTag(string, string) Telemetry

	Inc(name string, value float64, extra ...interface{}) Logger
	Gauge(string, float64, ...interface{}) Logger
	Fix(string, float64, ...interface{}) Logger
	Histogram(name string, value float64, extra ...interface{}) Logger
	Summary(name string, value float64, extra ...interface{}) Logger

	Clone() Telemetry
	SetLogger(l Logger)
}
