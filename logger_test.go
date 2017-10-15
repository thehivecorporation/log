package log

import (
	"testing"

	"github.com/cabify/fraud/assert"
)

func Test_removeRootFromGOPATH(t *testing.T) {
	result := removeRootFromPath("home/caster/go/src/github.com/thehivecorporation/log/test/main.go:20")
	expected := "github.com/thehivecorporation/log/test/main.go:20"

	assert.Equal(t, expected, result)
}
