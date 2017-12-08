# log
Another logger with telemetry capabilities using StatsD, Prometheus, InfluxDB...

## How to use it

You can use it like any common log package:

```go
import (
    "github.com/thehivecorporation/log/writers/text"
    "github.com/thehivecorporation/log"
)

log.SetWriter(text.New(os.Stdout))
log.SetLevel(log.LevelInfo)

log.Info("Some information")
log.Debug("Some Debug info")
```


## How to use telemetry:

Use any of the included implementations on telemetry folder:

```go
import (
    "github.com/thehivecorporation/log/telemetry/statsd"
    "github.com/thehivecorporation/log/writers/json"
    "github.com/thehivecorporation/log"
)

log.SetWriter(text.New(os.Stdout))
log.SetTelemetry(statsd.New(statsd.Conf{
    Address:   "localhost:9125",
    Namespace: "myapp.",
}))

log.WithTags("tag1").Inc("mycounter", 1).WithField("key", "value").Info("incremented")
// Outputs
// {"level":1,"messages":["incremented"],"fields":{"key":"value"},"ts":"2017-10-16T00:03:07.685786669+02:00"}
```

At the same time a +1 with tag 'tag1' to the metric 'mycounter'

## More information

Common functionality is covered by the following interface:

```go
type Tags map[string]string
type Fields map[string]interface{}
type Level int

type Logger interface {
	Debug(msg interface{}) Telemetry
	Info(msg interface{}) Telemetry
	Warn(msg interface{}) Telemetry
	Error(msg interface{}) Telemetry
	Fatal(msg interface{}) Telemetry

	Debugf(msg string, v ...interface{}) Telemetry
	Infof(msg string, v ...interface{}) Telemetry
	Warnf(msg string, v ...interface{}) Telemetry
	Errorf(msg string, v ...interface{}) Telemetry
	Fatalf(msg string, v ...interface{}) Telemetry

	WithField(s string, v interface{}) Logger
	WithFields(Fields) Logger
	WithError(...error) Logger
	WithErrors(...error) Logger

	WithTags(t Tags) Telemetry
	WithTag(string, string) Telemetry

	Clone(callStack int) Logger
}
```

Using **`WithTags(s ...string) Telemetry`** will return an instance to use it with the following methods:

```go
type Telemetry interface {
	WithTags(t Tags) Telemetry
	WithTag(string, string) Telemetry

	Inc(name string, value float64, extra ...interface{}) Logger
	Gauge(string, float64, ...interface{}) Logger
	Histogram(name string, value float64, extra ...interface{}) Logger

	Clone() Telemetry
	SetLogger(l Logger)
}
```

## Log outputs (Writers)

* **Text** with a provided io.Writer
* **JSON** with a provided io.Writer
* **Custom** will resend any output back for your control in case that you need to customize what to do with each output
* **Memory** (stores each log to an array, useful for testing purposes)
* **Multi**: Allows to use more than one Writer so you can log to console **AND** to Papertrail
* **Papertrail**
* TODO Kafka
* TODO Google PubSub
* TODO NATS
* TODO NSQ
* TODO AWS Kinesis
* TODO AWS SQS
* TODO Elastic

## Telemetry implementations:

* **StatsD**
* **Prometheus**
* **InfluxDB**
* **Multi** (alpha) to monitor using more than one system
* TODO OpenTSDB
Test
