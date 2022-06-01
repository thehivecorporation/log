package log

type dummyLogger struct {
	telemetry Telemetry
}

func (d dummyLogger) Debug(_ interface{}, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Info(_ interface{}, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Warn(_ interface{}, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Error(_ interface{}, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Fatal(_ interface{}, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Debugf(_ string, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Infof(_ string, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Warnf(_ string, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Errorf(_ string, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d dummyLogger) Fatalf(_ string, _ ...interface{}) Telemetry {
	return d.telemetry
}

func (d *dummyLogger) WithField(s string, v interface{}) Logger {
	return d
}

func (d *dummyLogger) WithFields(_ Fields) Logger {
	return d
}

func (d *dummyLogger) WithError(_ ...error) Logger {
	return d
}

func (d *dummyLogger) WithErrors(_ ...error) Logger {
	return d
}

func (d dummyLogger) WithTags(_ Tags) Telemetry {
	return d.telemetry
}

func (d dummyLogger) WithTag(_ string, _ string) Telemetry {
	return d.telemetry
}

func (d *dummyLogger) Clone() Logger {
	return d
}
