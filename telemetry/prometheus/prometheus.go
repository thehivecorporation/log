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
	summaries  map[string]*prometheus.SummaryVec
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

type Summaries []Summary

type Summary struct {
	Options prometheus.SummaryOpts
	Labels  []string
}

func New(c Counters, g Gauges, h Histograms, s Summaries) log.Telemetry {
	tel := telemetryImpl{
		counters:   make(map[string]*prometheus.CounterVec),
		histograms: make(map[string]*prometheus.HistogramVec),
		gauges:     make(map[string]*prometheus.GaugeVec),
		summaries:  make(map[string]*prometheus.SummaryVec),
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

	for _, v := range s {
		summary := prometheus.NewSummaryVec(v.Options, v.Labels)
		prometheus.MustRegister(summary)
		tel.summaries[v.Options.Name] = summary
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
	if m, ok := p.counters[name]; !ok {
		p.Logger.Errorf("Counter metric not found '%s'", name)
	} else {
		m.With(map[string]string(p.Tags)).Add(value)
	}

	return p.Logger
}

func (p *telemetryImpl) Gauge(name string, value float64, extra ...interface{}) log.Logger {
	if m, ok := p.gauges[name]; !ok {
		p.Logger.Errorf("Gauge metric not found '%s'", name)
	} else {
		m.With(map[string]string(p.Tags)).Set(value)
	}

	return p.Logger
}

func (p *telemetryImpl) Histogram(name string, value float64, extra ...interface{}) log.Logger {
	if m, ok := p.histograms[name]; !ok {
		p.Logger.Errorf("Histogram metric not found '%s'", name)
	} else {
		m.With(map[string]string(p.Tags)).Observe(value)
	}

	return p.Logger
}

func (p *telemetryImpl) Summary(name string, value float64, extra ...interface{}) log.Logger {
	if m, ok := p.summaries[name]; !ok {
		p.Logger.Errorf("Summary metric not found '%s'", name)
	} else {
		m.With(map[string]string(p.Tags)).Observe(value)
	}

	return p.Logger
}

func (p telemetryImpl) Clone() log.Telemetry {
	p.Tags = log.Tags{}

	return &p
}

func (p *telemetryImpl) SetLogger(l log.Logger) {
	p.Logger = l
}
