package log

import (
	"os"
	"time"
)

var loggerPrototype = logger{
	fields:       Fields{},
	telemetry:    &mockTelemetry{},
	start:        time.Now(),
	Writer:       newTextWriter(os.Stdout),
	level:        LevelInfo,
	includeStack: true,
}

var telemetryPrototype Telemetry = &mockTelemetry{
	Start: time.Now(),
}
