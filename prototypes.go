package log

import (
	"time"

	"github.com/fatih/color"
)

var loggerPrototype = logger{
	fields:       Fields{},
	telemetry:    &mockTelemetry{},
	start:        time.Now(),
	Writer:       newTextWriter(color.Output),
	level:        LevelInfo,
	includeStack: true,
}

var telemetryPrototype Telemetry = &mockTelemetry{
	Start: time.Now(),
}
