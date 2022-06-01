package log

type dummyLogger struct {
	telemetry Telemetry
}

func (d dummyLogger) Debug(msg interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Info(msg interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Warn(msg interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Error(msg interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Fatal(msg interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Debugf(msg string, v ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Infof(msg string, v ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Warnf(msg string, v ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Errorf(msg string, v ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Fatalf(msg string, v ...interface{}) Telemetry {
	return d.telemetry
}

func (d *dummyLogger) WithField(s string, v interface{}) Logger {
	return d
}

func (d *dummyLogger) WithFields(fields Fields) Logger {
	return d
}

func (d *dummyLogger) WithError(err ...error) Logger {
	return d
}

func (d *dummyLogger) WithErrors(err ...error) Logger {
	return d
}

func (d dummyLogger) WithTags(t Tags) Telemetry {
	return d.telemetry
}

func (d dummyLogger) WithTag(s string, s2 string) Telemetry {
	return d.telemetry
}

func (d *dummyLogger) Clone(callStack int) Logger {
	return d
}
