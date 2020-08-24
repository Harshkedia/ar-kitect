package main

import (
	"os"
	"testing"
)

func TestNothing(t *testing.T) {
	f, err := os.Stat("./web")
	t.Logf("%s %v", err, f)
	_ = f
	_ = err
}
