package log

func newTelemetry(l Logger) Telemetry {
	t := telemetryPrototype.Clone()
	t.SetLogger(l)

	return t
}

func newLogger() Logger {
	return loggerPrototype.Clone()
}

func newDummy(_ int) Logger {
	return &dummyLogger{}
}
