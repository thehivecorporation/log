package log

func Info(s string) Telemetry {
	return newLogger().Info(s)
}

func Debug(s string) Telemetry {
	return newLogger().Debug(s)
}

func Error(s string) Telemetry {
	return newLogger().Error(s)
}

func Warn(s string) Telemetry {
	return newLogger().Warn(s)
}

func Fatal(s string) Telemetry {
	return newLogger().Fatal(s)
}

func Infof(s string) Telemetry {
	return newLogger().Infof(s)
}

func Debugf(s string) Telemetry {
	return newLogger().Debugf(s)
}

func Errorf(s string) Telemetry {
	return newLogger().Errorf(s)
}

func Warnf(s string) Telemetry {
	return newLogger().Warnf(s)
}

func Fatalf(s string) Telemetry {
	return newLogger().Fatalf(s)
}

func WithFields(f Fields) Logger {
	return newLogger().WithFields(f)
}

func WithField(s string, v interface{}) Logger {
	return newLogger().WithField(s, v)
}

func WithTags(s ...string) Telemetry {
	return newTelemetry(newLogger()).WithTags(s...)
}

func Inc(name string, value float64) Logger {
	l := newLogger()
	newTelemetry(l).Inc(name, value)
	return l
}

func WithError(err error) Logger {
	return newLogger().WithError(err)
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
