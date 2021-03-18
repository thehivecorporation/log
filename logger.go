package log

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	juju_err "github.com/juju/errors"
)

type logger struct {
	fields    Fields
	telemetry Telemetry
	start     time.Time
	errors    []string
	Writer
	level        Level
	includeStack bool
}

func (l logger) WithErrors(errs ...error) Logger {
	return l.WithError(errs...)
}

func (l logger) WithError(errs ...error) Logger {
	if len(errs) == 1 {
		l.fillErrors("error", errs[0])
	} else {
		for i, err := range errs {
			l.fillErrors(fmt.Sprintf("error-%d", i), err)
		}
	}

	return &l
}

func (l *logger) writeLog(msg interface{}, level Level, more ...interface{}) Telemetry {
	var startLine string
	var file string
	var line int

	if l.includeStack {
		var i int
		for {
			_, file, line, _ = runtime.Caller(i)
			if !strings.Contains(file, "github.com/thehivecorporation/log"){
				_, file, line, _ = runtime.Caller(i)
				startLine = fmt.Sprintf("%s:%d", removeRootFromPath(file), line) + " %v"
				break
			}
			i++
		}
	} else {
		startLine = "%v"
	}

	//startLine = fmt.Sprintf("%s:%d", removeRootFromPath(file), line) + " %v"

	if len(more) > 0 && more[0] != nil {
		for i := 0; i < len(more); i++ {
			more[i] = fmt.Sprintf(" %v", more[i])
		}
		return l.checkLevelAndWriteLog(fmt.Sprintf(startLine, fmt.Sprint(append([]interface{}{msg}, more...)...)), level)
	} else {
		return l.checkLevelAndWriteLog(fmt.Sprintf(startLine, msg), level)
	}
}

func (l *logger) Debug(msg interface{}, more ...interface{}) Telemetry {
	return l.writeLog(msg, LevelDebug, more...)
}

func (l *logger) Warn(msg interface{}, more ...interface{}) Telemetry {
	return l.writeLog(msg, LevelWarn, more...)
}

func (l *logger) Fatal(msg interface{}, more ...interface{}) Telemetry {
	l.writeLog(msg, LevelFatal, more...)

	os.Exit(1)
	return nil
}

func (l *logger) Error(msg interface{}, more ...interface{}) Telemetry {
	return l.writeLog(msg, LevelError, more...)
}

func (l *logger) writeLogf(msg string, level Level, v ...interface{}) Telemetry {
	var i int
	var file string
	var line int
	for {
		_, file, line, _ = runtime.Caller(i)
		if !strings.Contains(file, "github.com/thehivecorporation/log"){
			_, file, line, _ = runtime.Caller(i)
			break
		}
		i++
	}
	startLine := fmt.Sprintf("%s:%d", removeRootFromPath(file), line)
	startLine += " %v"
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), level)
}

func (l *logger) Debugf(msg string, v ...interface{}) Telemetry {
	return l.writeLogf(msg, LevelDebug, v...)
}

func (l *logger) Infof(msg string, v ...interface{}) Telemetry {
	return l.writeLogf(msg, LevelInfo, v...)
}

func (l *logger) Errorf(msg string, v ...interface{}) Telemetry {
	return l.writeLogf(msg, LevelError, v...)
}

func (l *logger) Warnf(msg string, v ...interface{}) Telemetry {
	return l.writeLogf(msg, LevelWarn, v...)
}

func (l *logger) Fatalf(msg string, v ...interface{}) Telemetry {
	return l.writeLogf(msg, LevelFatal, v...)

	os.Exit(1)

	return l.telemetry
}

func (l *logger) Info(s interface{}, more ...interface{}) Telemetry {
	return l.writeLog(s, LevelInfo, more...)
}

func (l logger) Clone() Logger {
	l.fields = Fields{}
	l.errors = make([]string, 0)
	l.telemetry = telemetryPrototype

	return &l
}

func (l *logger) WithFields(f Fields) Logger {
	l.fields = f
	return l
}

func (l logger) WithField(s string, v interface{}) Logger {
	l.fields[s] = v

	return &l
}

func (l *logger) WithTag(k string, v string) Telemetry {
	return newTelemetry(l).WithTags(Tags{k: v})
}

func (l *logger) WithTags(tags Tags) Telemetry {
	return newTelemetry(l).WithTags(tags)
}

func (l *logger) checkLevelAndWriteLog(msg interface{}, level Level) Telemetry {
	if l.level > level {
		return l.telemetry
	}

	l.writeLogCommon(msg, level)

	return l.telemetry
}

func (l *logger) writeLogCommon(msg interface{}, level Level) {
	l.WriteLog(&Payload{
		Messages:          []interface{}{msg},
		ElapsedSinceStart: time.Since(l.start) / time.Second,
		Timestamp:         time.Now(),
		Level:             level,
		Fields:            l.fields,
		Errors:            l.errors,
	})
}

var root = regexp.MustCompile(".*?go/src/(.*$)")

func removeRootFromPath(s string) string {
	sm := root.FindStringSubmatch(s)
	if len(sm) < 2 {
		return s
	}

	return sm[1]
}

func (l *logger) fillErrors(mapkey string, err error) {
	switch e := err.(type) {
	case *juju_err.Err:
		l.errors = make([]string, 0)
		stacktrace := e.StackTrace()

		for i := 0; i < len(stacktrace); i++ {
			l.errors = append(l.errors, removeRootFromPath(stacktrace[i]))
		}
	default:
		l.fields[mapkey] = e.Error()
	}
}
