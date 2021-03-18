package log

import (
	"time"
)

type mockTelemetry struct {
	Logger Logger
	Tags   Tags
	Start  time.Time
}

func (m *mockTelemetry) Gauge(_ string, _ float64, _ ...interface{}) Logger {
	return m.Logger
}

func (m *mockTelemetry) Fix(_ string, _ float64, _ ...interface{}) Logger {
	return m.Logger
}

func (m *mockTelemetry) Histogram(_ string, _ float64, _ ...interface{}) Logger {
	return m.Logger
}

func (m *mockTelemetry) SetLogger(l Logger) {
	m.Logger = l
}

func (m *mockTelemetry) Inc(_ string, _ float64, _ ...interface{}) Logger {
	return m.Logger
}

func (m *mockTelemetry) Summary(_ string, _ float64, _ ...interface{}) Logger {
	return m.Logger
}

func (m *mockTelemetry) WithTag(k, v string) Telemetry {
	//TODO Remove this check. It should not be necessary
	if m.Tags == nil {
	}
	return m
}

func (m *mockTelemetry) WithTags(t Tags) Telemetry {
	return m
}

func (m mockTelemetry) Clone() Telemetry {
	m.Start = time.Now()
	return &m
}
