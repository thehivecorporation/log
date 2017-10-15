package log

import (
	"fmt"
	"time"
)

type mockTelemetry struct {
	Logger Logger
	Tags   []string
	Start  time.Time
}

func (m *mockTelemetry) Gauge(_ string, _ float64, _ ...interface{}) Logger {
	fmt.Println("GAUGE", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) Histogram(_ string, _ float64, _ ...interface{}) Logger {
	fmt.Println("HIST", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) SetLogger(l Logger) {
	m.Logger = l
}

func (m *mockTelemetry) Inc(_ string, _ float64, _ ...interface{}) Logger {
	fmt.Println("INC", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) WithTags(s ...string) Telemetry {
	m.Tags = append(m.Tags, s...)
	return m
}

func (m mockTelemetry) Clone() Telemetry {
	m.Tags = make([]string, 0)
	return &m
}
