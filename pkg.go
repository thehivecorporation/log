package log

func Info(s string) Telemetry {
	return newLogger(2).Info(s)
}

func Debug(s string) Telemetry {
	return newLogger(2).Debug(s)
}

func Error(s string) Telemetry {
	return newLogger(2).Error(s)
}

func Warn(s string) Telemetry {
	return newLogger(2).Warn(s)
}

func Fatal(s string) Telemetry {
	return newLogger(2).Fatal(s)
}

func Infof(s string) Telemetry {
	return newLogger(2).Infof(s)
}

func Debugf(s string) Telemetry {
	return newLogger(2).Debugf(s)
}

func Errorf(s string) Telemetry {
	return newLogger(2).Errorf(s)
}

func Warnf(s string) Telemetry {
	return newLogger(2).Warnf(s)
}

func Fatalf(s string) Telemetry {
	return newLogger(2).Fatalf(s)
}

func WithFields(f Fields) Logger {
	return newLogger(2).WithFields(f)
}

func WithField(s string, v interface{}) Logger {
	return newLogger(2).WithField(s, v)
}

func WithTags(s ...string) Telemetry {
	return newTelemetry(newLogger(1)).WithTags(s...)
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
	loggerPrototype.telemetry = t
	telemetryPrototype = t
}
