package multi

import (
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"
)

type telemetryImpl struct {
	telemetry.Common
	impls []log.Telemetry
}

func New(ts ...log.Telemetry) log.Telemetry {
	return &telemetryImpl{
		impls: ts,
	}
}

func (t *telemetryImpl) WithTags(s log.Tags) log.Telemetry {
	for _, i := range t.impls {
		i.WithTags(s)
	}

	return t
}

func (t *telemetryImpl) WithTag(k, v string) log.Telemetry {
	for _, i := range t.impls {
		i.WithTag(k, v)
	}

	return t
}

func (t *telemetryImpl) Inc(name string, value float64, extra ...interface{}) log.Logger {
	for _, i := range t.impls {
		i.Inc(name, value, extra)
	}

	return t.Logger
}

func (t *telemetryImpl) Gauge(name string, value float64, extra ...interface{}) log.Logger {
	for _, i := range t.impls {
		i.Gauge(name, value, extra)
	}

	return t.Logger
}

func (t *telemetryImpl) Fix(name string, value float64, extra ...interface{}) log.Logger {
	for _, i := range t.impls {
		i.Fix(name, value, extra)
	}

	return t.Logger
}

func (t *telemetryImpl) Histogram(name string, value float64, extra ...interface{}) log.Logger {
	for _, i := range t.impls {
		i.Histogram(name, value, extra)
	}

	return t.Logger
}

func (t *telemetryImpl) Summary(name string, value float64, extra ...interface{}) log.Logger {
	for _, i := range t.impls {
		i.Summary(name, value, extra)
	}

	return t.Logger
}

func (t telemetryImpl) Clone() log.Telemetry {
	t.impls = make([]log.Telemetry, len(t.impls))

	for i, v := range t.impls {
		t.impls[i] = v.Clone()
	}

	t.Tags = make(map[string]string, 0)

	return &t
}

func (t *telemetryImpl) SetLogger(l log.Logger) {
	t.Logger = l
}
