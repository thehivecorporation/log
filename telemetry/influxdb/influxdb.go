package influxdb

import (
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/telemetry"

	"github.com/influxdata/influxdb/client/v2"
	"time"
)

func New(c client.HTTPConfig, points []client.BatchPointsConfig) log.Telemetry {
	cli, err := client.NewHTTPClient(c)
	if err != nil {
		log.WithError(err).Fatal("Could not create client to InfluxDB")
	}

	t := &telemetryImpl{
		c:      cli,
		points: make(map[string]client.BatchPoints),
		Common: telemetry.Common{
			Tags:  make([]string, 0),
			Start: time.Now(),
		},
		influxTags: make(map[string]string),
	}

	if len(points) == 0 {
		log.Fatal("No points attached")
	}

	for _, p := range points {
		if t.points[p.Database], err = client.NewBatchPoints(p); err != nil {
			log.WithError(err).Fatal("Error creating batch point")
		}
	}

	return t
}

type telemetryImpl struct {
	telemetry.Common
	c          client.Client
	points     map[string]client.BatchPoints
	influxTags map[string]string
}

func (t *telemetryImpl) WithTags(s ...string) log.Telemetry {
	for i := 0; i < len(s)-1; i += 2 {
		t.influxTags[s[i]] = s[i+1]
	}

	return t
}

func (t *telemetryImpl) Inc(name string, value float64, extras ...interface{}) log.Logger {
	return t.addPoint(name, value, extras)
}

func (t *telemetryImpl) addPoint(name string, value float64, extras ...interface{}) log.Logger {
	var batchPoint client.BatchPoints
	var fields map[string]interface{}

	for _, extra := range extras {
		switch type_ := extra.(type) {
		case string:
			if t.points[type_] == nil {
				log.WithField("batch_point_name", type_).Error("Logging error: Batch point not found")
				return t.Logger
			}

			batchPoint = t.points[type_]
		case map[string]interface{}:
			fields = type_
		}
	}

	if fields != nil && batchPoint != nil {
		if p, err := client.NewPoint(name, t.influxTags, fields); err != nil {
			log.WithError(err).Error("Could not create point")
		} else {
			batchPoint.AddPoint(p)
		}
	}

	return t.Logger
}

func (t *telemetryImpl) Gauge(n string, v float64, extra ...interface{}) log.Logger {
	return t.addPoint(n, v, extra)
}

func (t *telemetryImpl) Histogram(name string, value float64, extra ...interface{}) log.Logger {
	return t.addPoint(name, value, extra)
}

func (t telemetryImpl) Clone() log.Telemetry {
	t.Tags = log.Tags{}
	t.influxTags = make(map[string]string)

	return &t
}

func (t *telemetryImpl) SetLogger(l log.Logger) {
	t.Logger = l
}
