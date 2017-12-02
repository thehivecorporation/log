package log

func Info(s interface{}) Telemetry {
	return newLogger(2).Info(s)
}

func Debug(s interface{}) Telemetry {
	return newLogger(2).Debug(s)
}

func Error(s interface{}) Telemetry {
	return newLogger(2).Error(s)
}

func Warn(s interface{}) Telemetry {
	return newLogger(2).Warn(s)
}

func Fatal(s interface{}) Telemetry {
	return newLogger(2).Fatal(s)
}

func Infof(s string, i ...interface{}) Telemetry {
	return newLogger(2).Infof(s, i...)
}

func Debugf(s string, i ...interface{}) Telemetry {
	return newLogger(2).Debugf(s, i...)
}

func Errorf(s string, i ...interface{}) Telemetry {
	return newLogger(2).Errorf(s, i...)
}

func Warnf(s string, i ...interface{}) Telemetry {
	return newLogger(2).Warnf(s, i...)
}

func Fatalf(s string, i ...interface{}) Telemetry {
	return newLogger(2).Fatalf(s, i...)
}

func WithFields(f Fields) Logger {
	return newLogger(2).WithFields(f)
}

func WithField(s string, v interface{}) Logger {
	return newLogger(2).WithField(s, v)
}

func WithTag(k string, v string) Telemetry {
	return newTelemetry(newLogger(1)).WithTag(k, v)
}

func WithTags(tags Tags) Telemetry {
	return newTelemetry(newLogger(1)).WithTags(tags)
}

func Inc(name string, value float64) Logger {
	l := newLogger(1)
	newTelemetry(l).Inc(name, value)
	return l
}

func WithError(err ...error) Logger {
	return newLogger(1).WithError(err...)
}

func SetLevel(l Level) {
	loggerPrototype.level = l
}

func SetWriter(w Writer) {
	loggerPrototype.Writer = w
}

func SetTelemetry(t Telemetry) {
	t.SetLogger(&loggerPrototype)
	loggerPrototype.telemetry = t
	telemetryPrototype = t
}
