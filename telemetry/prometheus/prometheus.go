package log_prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"
)

type telemetryImpl struct {
	telemetry.Common

	counters   map[string]*prometheus.CounterVec
	gauges     map[string]*prometheus.GaugeVec
	histograms map[string]*prometheus.HistogramVec
}

type Opt struct {
	Options prometheus.Opts
	Labels  []string
}
type Opts []Opt

type Counters Opts
type Gauges Opts

type Histograms []Histogram
type Histogram struct {
	Options prometheus.HistogramOpts
	Labels  []string
}

func New(c Counters, g Gauges, h Histograms) log.Telemetry {
	tel := telemetryImpl{
		counters:   make(map[string]*prometheus.CounterVec),
		histograms: make(map[string]*prometheus.HistogramVec),
		gauges:     make(map[string]*prometheus.GaugeVec),
	}

	for _, v := range c {
		gauge := prometheus.NewCounterVec(prometheus.CounterOpts(v.Options), v.Labels)
		prometheus.MustRegister(gauge)
		tel.counters[v.Options.Name] = gauge
	}

	for _, v := range g {
		counter := prometheus.NewGaugeVec(prometheus.GaugeOpts(v.Options), v.Labels)
		prometheus.MustRegister(counter)
		tel.gauges[v.Options.Name] = counter
	}

	for _, v := range h {
		histogram := prometheus.NewHistogramVec(v.Options, v.Labels)
		prometheus.MustRegister(histogram)
		tel.histograms[v.Options.Name] = histogram
	}

	return &tel
}

func (p *telemetryImpl) WithTags(t log.Tags) log.Telemetry {
	p.Tags = t

	return p
}

func (p *telemetryImpl) WithTag(k string, v string) log.Telemetry {
	p.Tags[k] = v

	return p
}

func (p *telemetryImpl) Inc(name string, value float64, extra ...interface{}) log.Logger {
	if p.counters[name] == nil {
		p.Logger.Errorf("Counter metric not found '%s'", name)
		return p.Logger
	}

	p.counters[name].With(map[string]string(p.Tags)).Add(value)

	return p.Logger
}

func (p *telemetryImpl) Gauge(name string, value float64, extra ...interface{}) log.Logger {
	if p.gauges[name] == nil {
		p.Logger.Errorf("Gauge metric not found '%s'", name)
		return p.Logger
	}

	p.gauges[name].With(map[string]string(p.Tags)).Set(value)

	return p.Logger
}

func (p *telemetryImpl) Histogram(name string, value float64, extra ...interface{}) log.Logger {
	if p.histograms[name] == nil {
		p.Logger.Errorf("Histogram metric not found '%s'", name)
		return p.Logger
	}

	p.histograms[name].With(map[string]string(p.Tags)).Observe(value)

	return p.Logger
}

func (p telemetryImpl) Clone() log.Telemetry {
	p.Tags = log.Tags{}

	return &p
}

func (p *telemetryImpl) SetLogger(l log.Logger) {
	p.Logger = l
}
