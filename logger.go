package log

import (
	"fmt"
	"os"
	"runtime"
	"time"

	juju_err "github.com/juju/errors"
	"strings"
)

type logger struct {
	callStack int
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

func (l *logger) Debug(msg interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, msg), LevelDebug)
}

func (l *logger) Warn(msg interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, msg), LevelWarn)
}

func (l *logger) Fatal(msg interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, msg), LevelFatal)

	os.Exit(1)

	return nil
}

func (l *logger) Error(msg interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, msg), LevelError)
}

func (l *logger) Debugf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelDebug)
}

func (l *logger) Infof(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelInfo)
}

func (l *logger) Errorf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelError)
}

func (l *logger) Warnf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	return l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelWarn)
}

func (l *logger) Fatalf(msg string, v ...interface{}) Telemetry {
	_, file, line, _ := runtime.Caller(l.callStack)
	l.checkLevelAndWriteLog(fmt.Sprintf("%s:%d %v", removeRootFromPath(file), line, fmt.Sprintf(msg, v...)), LevelFatal)

	os.Exit(1)

	return l.telemetry
}

func (l *logger) Info(s interface{}) Telemetry {
	return l.checkLevelAndWriteLog(s, LevelInfo)
}

func (l logger) Clone(callStack int) Logger {
	l.fields = Fields{}
	l.errors = make([]string, 0)
	l.telemetry = telemetryPrototype
	l.callStack = callStack
	l.start = time.Now()

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

func (l *logger) errJujuStack(msg interface{}) Telemetry {
	l.WriteLog(&Payload{
		Messages:  []interface{}{msg},
		Timestamp: &l.start,
		Level:     LevelError,
		Fields:    l.fields,
	})

	return l.telemetry
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
		Messages:  []interface{}{msg},
		Timestamp: &l.start,
		Level:     level,
		Fields:    l.fields,
		Errors:    l.errors,
	})
}

func removeRootFromPath(s string) string {
	i := strings.Index(s, "go/src")

	return s[i+7:]
}

func (l *logger) fillErrors(mapkey string, err error) {
	switch e := err.(type) {
	case *juju_err.Err:
		l.errors = make([]string, len(e.StackTrace()))
		for i, subErr := range e.StackTrace() {
			l.errors[i] = removeRootFromPath(subErr)
		}
	default:
		l.fields[mapkey] = e.Error()
	}
}
