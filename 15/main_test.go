package main

import (
	"os"
	"testing"
)

// ---- Test Setup ------------------------------

func TestMain(m *testing.M) {

	exitCode := m.Run()
	os.Exit(exitCode)
}

// ---- Tests -----------------------------------
