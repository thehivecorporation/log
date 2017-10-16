package log

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"strings"

	juju_err "github.com/juju/errors"
)

type logger struct {
	fields    Fields
	telemetry Telemetry
	start     time.Time
	errors    []string
	Writer
	level Level
}

func (l logger) WithErrors(errs ...error) Logger {
	return l.WithError(errs...)
}

func (l logger) fillErrors(mapkey string, err error) {
	switch e := err.(type) {
	case *juju_err.Err:
		l.errors = e.StackTrace()
	default:
		l.fields[mapkey] = removeRootFromPath(e.Error())
	}
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

func (l *logger) Debug(msg string) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, msg), LevelDebug)
}

func (l *logger) Warn(msg string) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, msg), LevelWarn)
}

func (l *logger) Fatal(msg string) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, msg), LevelFatal)

	os.Exit(1)

	return nil
}

func (l *logger) Debugf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelDebug)
}

func (l *logger) Infof(msg string, v ...interface{}) Telemetry {
	return l.checkLevelAndWriteLog(fmt.Sprintf(msg, v...), LevelInfo)
}

func (l *logger) Warnf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelWarn)
}

func (l *logger) Error(msg string) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, msg), LevelError)
}

func (l *logger) Errorf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelError)
}

func (l *logger) Fatalf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(1)
	l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %s", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelFatal)

	os.Exit(1)

	return l.telemetry
}

func (l *logger) Info(s string) Telemetry {
	return l.checkLevelAndWriteLog(s, LevelInfo)
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

func (l *logger) WithTags(s ...string) Telemetry {
	return newTelemetry(l).WithTags(s...)
}

func (l *logger) errJujuStack(msg string) Telemetry {
	l.WriteLog(&Payload{
		Messages:  []string{msg},
		Timestamp: &l.start,
		Level:     LevelError,
		Fields:    l.fields,
	})

	return l.telemetry
}

func (l *logger) checkLevelAndWriteLog(msg string, level Level) Telemetry {
	if l.level > level {
		return l.telemetry
	}

	l.writeLogCommon(msg, level)

	return l.telemetry
}

func (l *logger) writeLogCommon(msg string, level Level) {
	l.WriteLog(&Payload{
		Messages:  []string{msg},
		Timestamp: &l.start,
		Level:     level,
		Fields:    l.fields,
		errors:    l.errors,
	})
}

func removeRootFromPath(s string) string {
	i := strings.Index(s, "go/src")

	return s[i+7:]
}
