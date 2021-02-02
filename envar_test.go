package envar

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	Load()

	if os.Getenv("FOO") != "BAR" {
		t.Fail()
	}
	if os.Getenv("GO") != "WALK" {
		t.Fail()
	}
}
