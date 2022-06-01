package log

func Info(s interface{}, more ...interface{}) Telemetry {
	return newLogger().Info(s)
}

func Debug(s interface{}, more ...interface{}) Telemetry {
	return newLogger().Debug(s, more...)
}

func Error(s interface{}, more ...interface{}) Telemetry {
	return newLogger().Error(s, more...)
}

func Warn(s interface{}, more ...interface{}) Telemetry {
	return newLogger().Warn(s, more...)
}

func Fatal(s interface{}, more ...interface{}) Telemetry {
	return newLogger().Fatal(s, more...)
}

func FatalIfError(err error) Telemetry {
	if err == nil {
		return newTelemetry(newDummy(2))
	}

	return newLogger().Fatal(err.Error())
}

func FatalIfErrorS(err error, s interface{}) Telemetry {
	if err == nil {
		return newTelemetry(newDummy(2))
	}
	return newLogger().Fatal(s)
}

func FatalFIfError(err error, s string, i ...interface{}) Telemetry {
	if err == nil {
		return newTelemetry(newDummy(2))
	}

	return newLogger().Fatalf(s, i...)
}

func Infof(s string, i ...interface{}) Telemetry {
	return newLogger().Infof(s, i...)
}

func Debugf(s string, i ...interface{}) Telemetry {
	return newLogger().Debugf(s, i...)
}

func Errorf(s string, i ...interface{}) Telemetry {
	return newLogger().Errorf(s, i...)
}

func Warnf(s string, i ...interface{}) Telemetry {
	return newLogger().Warnf(s, i...)
}

func Fatalf(s string, i ...interface{}) Telemetry {
	return newLogger().Fatalf(s, i...)
}

func WithFields(f Fields) Logger {
	return newLogger().WithFields(f)
}

func WithField(s string, v interface{}) Logger {
	return newLogger().WithField(s, v)
}

func WithTag(k string, v string) Telemetry {
	return newTelemetry(newLogger()).WithTag(k, v)
}

func WithTags(tags Tags) Telemetry {
	return newTelemetry(newLogger()).WithTags(tags)
}

func Histogram(name string, value float64, extra ...interface{}) Logger {
	l := newLogger()
	newTelemetry(l).Histogram(name, value, extra)
	return l
}

func Summary(name string, value float64, extra ...interface{}) Logger {
	l := newLogger()
	newTelemetry(l).Summary(name, value, extra)
	return l
}

func Inc(name string, value float64) Logger {
	l := newLogger()
	newTelemetry(l).Inc(name, value)
	return l
}

func Gauge(name string, value float64) Logger {
	l := newLogger()
	newTelemetry(l).Gauge(name, value)
	return l
}

func Fix(name string, value float64) Logger {
	l := newLogger()
	newTelemetry(l).Fix(name, value)
	return l
}

func WithError(err ...error) Logger {
	return newLogger().WithError(err...)
}

func SetLevel(l Level) {
	loggerPrototype.level = l
}

func SetWriter(w Writer) {
	loggerPrototype.Writer = w
}

func DisableStackInfo() {
	loggerPrototype.includeStack = false
}

func SetTelemetry(t Telemetry) {
	t.SetLogger(&loggerPrototype)
	loggerPrototype.telemetry = t
	telemetryPrototype = t
}
