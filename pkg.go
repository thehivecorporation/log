package log

func Info(s interface{}, more ...interface{}) Telemetry {
	return newLogger(2).Info(s)
}

func Debug(s interface{}, more ...interface{}) Telemetry {
	return newLogger(2).Debug(s, more...)
}

func Error(s interface{}, more ...interface{}) Telemetry {
	return newLogger(2).Error(s, more...)
}

func Warn(s interface{}, more ...interface{}) Telemetry {
	return newLogger(2).Warn(s, more...)
}

func Fatal(s interface{}, more ...interface{}) Telemetry {
	return newLogger(2).Fatal(s, more...)
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

func Histogram(name string, value float64, extra ...interface{}) Logger {
	l := newLogger(1)
	newTelemetry(l).Histogram(name, value, extra)
	return l
}

func Summary(name string, value float64, extra ...interface{}) Logger {
	l := newLogger(1)
	newTelemetry(l).Summary(name, value, extra)
	return l
}

func Inc(name string, value float64) Logger {
	l := newLogger(1)
	newTelemetry(l).Inc(name, value)
	return l
}

func Gauge(name string, value float64) Logger {
	l := newLogger(1)
	newTelemetry(l).Gauge(name, value)
	return l
}

func Fix(name string, value float64) Logger {
	l := newLogger(1)
	newTelemetry(l).Fix(name, value)
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

func DisableStackInfo() {
	loggerPrototype.includeStack = false
}

func SetTelemetry(t Telemetry) {
	t.SetLogger(&loggerPrototype)
	loggerPrototype.telemetry = t
	telemetryPrototype = t
}
