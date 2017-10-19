package statsd

import (
	"fmt"
	statsd_client "github.com/DataDog/datadog-go/statsd"
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"
)

type telemetryImpl struct {
	telemetry.Common
	c *statsd_client.Client
}

type Conf struct {
	Address   string
	Namespace string
	Tags      []string
}

func New(conf Conf) log.Telemetry {
	c, err := statsd_client.New(conf.Address)
	if err != nil {
		log.WithError(err).WithField("address", conf.Address).Fatal("Could not open statsD client")
	}

	c.Namespace = conf.Namespace
	c.Tags = conf.Tags

	return &telemetryImpl{
		c: c,
	}
}

func (s *telemetryImpl) WithTag(k, v string) log.Telemetry {
	s.Tags[k] = v
	return s
}

func (s *telemetryImpl) WithTags(tags log.Tags) log.Telemetry {
	s.Tags = tags
	return s
}

func (s *telemetryImpl) Inc(name string, value float64, extra ...interface{}) log.Logger {
	if err := s.c.Incr(name, s.getTagsAr(), value); err != nil {
		s.Logger.WithError(err).WithFields(log.Fields{"tags": s.Tags}).Errorf("Error trying to increment metric '%s'", name)
	}

	return s.Logger
}

func (s *telemetryImpl) Gauge(name string, value float64, extra ...interface{}) log.Logger {
	if err := s.c.Gauge(name, value, s.getTagsAr(), 1); err != nil {
		s.Logger.Error(err.Error())
	}

	return s.Logger
}

func (s *telemetryImpl) Histogram(name string, value float64, extra ...interface{}) log.Logger {
	if err := s.c.Histogram(name, value, s.getTagsAr(), 1); err != nil {
		s.Logger.Error(err.Error())
	}

	return s.Logger
}

func (s telemetryImpl) Clone() log.Telemetry {
	s.Tags = log.Tags{}
	return &s
}

func (s *telemetryImpl) SetLogger(l log.Logger) {
	s.Logger = l
}

func (s *telemetryImpl) getTagsAr() (tagsAr []string) {
	tagsAr = make([]string, len(s.Tags))
	var i int
	for k, v := range s.Tags {
		tagsAr[i] = fmt.Sprintf("%s:%s", k, v)
		i++
	}

	return tagsAr
}
