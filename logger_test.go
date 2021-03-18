package log

import (
	"github.com/juju/errors"
	"testing"
)

func TestRemoveRootFromPath(t *testing.T){
	res := removeRootFromPath("go/src/github.com/thehivecorporation/log")
	if res != "github.com/thehivecorporation/log" {
		t.Fail()
	}

	WithError(errors.Annotate(stacktrace(), "another")).Error("Error")

	res = removeRootFromPath("something_else")
	if res != "something_else" {
		t.Fail()
	}

	t.Log(res)
}

func stacktrace()error{
	return errors.New("stacktrace")
}