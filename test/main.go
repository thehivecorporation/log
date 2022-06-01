package main

import (
	"os"

	"github.com/juju/errors"
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/writers/json"
)

func main() {
	log.WithField("hello", "world").Info("Hello from external/main.go")

	log.WithTags(log.Tags{"endpoint": "e1", "host": "h1"}).
		Inc("counter", 1).
		WithField("field", 2).
		Info("Hello")

	log.WithField("key1", 3).
		WithTag("endpoint", "e2").
		WithTag("endpoint", "e3").
		Inc("counter", 1).
		Info("Uf")

	log.WithField("key2", "value").WithTag("db", "tag5").WithTag("endpoint", "tag6").Inc("counter3", 11).Info("Uf2")
	log.Inc("counter3", 11).Info("Uf2")
	log.WithField("errorkey", "errorvalue").Error("An error")
	//log.WithField("errorkey", "errorvalue").WithError(po_error.New("My error")).Error("Another error")
	log.WithField("database", "users").WithTag("database", "db1").Inc("counter", 1)

	err := errors.New("An error")
	err = errors.Annotatef(err, "A wrapper")

	log.WithError(err).Error("Error chungo")

	log.WithField("key", "value").Debug("Debug message")

	log.SetWriter(json.New(os.Stdout))
	log.Warn("Warn message")

	err = errors.New("An error")
	err = errors.Annotatef(err, "A wrapper")
	log.WithError(err).Error("Error chungo")

	err = errors.Errorf("%d errors", 10)
	err = errors.Annotatef(err, "A wrapper")
	log.WithError(err).Error("MANY ERRORS")
	//
	//log.SetWriter(text.New(os.Stdout))
	//log.Info("Changing level to info")
	//log.SetLevel(log.LevelInfo)
	//log.Warn("You should be reading this warn message")
	//log.Error("You should be reading this error message")
	log.Info("You should be reading this info message")
	log.Debug("You shouldn't be reading this debug message")

	log.Info("Changing level to error")
	log.SetLevel(log.LevelError)
	log.Warn("You should be reading this warn message")
	log.Error("You should be reading this error message")
	log.Info("You should be reading this info message")
	log.Debug("You shouldn't be reading this debug message")

	//Now with statsd
	log.SetLevel(log.LevelDebug)
	log.SetWriter(json.New(os.Stdout))
	///FIXME
	//log.SetTelemetry(statsd.New(statsd.Conf{
	//	Address:   "localhost:9125",
	//	Namespace: "myapp.",
	//}))

	log.WithField("key", "value").WithTag("endpoint", "e4").Inc("mycounter", 1).Info("incremented")

	log.FatalIfError(nil)
	log.FatalIfErrorS(nil, "no fatal")
	//log.FatalIfError(errors.New("error"), "fatal")
}
