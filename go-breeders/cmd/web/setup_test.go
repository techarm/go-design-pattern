package main

import (
	"go-breeders/adapters"
	"go-breeders/config"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}
	testApp = application{
		App: config.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}
