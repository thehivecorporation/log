package log

func newTelemetry(l Logger) Telemetry {
	t := telemetryPrototype.Clone()
	t.SetLogger(l)

	return t
}

func newLogger(callStack int) Logger {
	return loggerPrototype.Clone(callStack)
}
