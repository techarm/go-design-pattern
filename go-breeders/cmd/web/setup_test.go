package main

import (
	"go-breeders/config"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: config.New(nil),
	}

	os.Exit(m.Run())
}
