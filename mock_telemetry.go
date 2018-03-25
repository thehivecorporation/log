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
	//fmt.Println("GAUGE", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) Histogram(_ string, _ float64, _ ...interface{}) Logger {
	//fmt.Println("HIST", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) SetLogger(l Logger) {
	m.Logger = l
}

func (m *mockTelemetry) Inc(_ string, _ float64, _ ...interface{}) Logger {
	//fmt.Println("INC", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) Summary(_ string, _ float64, _ ...interface{}) Logger {
	//fmt.Println("INC", m.Tags)
	return m.Logger
}

func (m *mockTelemetry) WithTag(k, v string) Telemetry {
	//TODO Remove this check. It should not be necessary
	if m.Tags == nil {
		//m.Tags = make(map[string]string)
	}
	//m.Tags[k] = v
	return m
}

func (m *mockTelemetry) WithTags(t Tags) Telemetry {
	//m.Tags = t
	return m
}

func (m mockTelemetry) Clone() Telemetry {
	//m.Tags = make(map[string]string)
	m.Start = time.Now()
	return &m
}
