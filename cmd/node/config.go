package main

import (
	"flag"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var configuration Config
var neighborSearchTimeout bool
var neighborSearchTimeoutExpiration time.Time

// https://stackoverflow.com/questions/35809252/check-if-flag-was-provided-in-go
func confirmServerValueProvided() bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "server" {
			found = true
		}
	})
	return found
}

func validateServerAddress(address string) bool {
	match, err := regexp.MatchString("^http://[0-9]\\.[0-9]\\.[0-9]\\.[0-9]:[0-9]+$", address)
	if err != nil {
		panic(fmt.Errorf(fmt.Sprintf("Invalid server address: %s", address)))
	}
	if !match {
		panic(fmt.Errorf(fmt.Sprintf("Invalid server address: %s", address)))
	}
	return true
}

// If no server is provided create a baseline configuration and wait for more nodes to connect and establish web.
func bootstrapSelf() {
	var e error
	neighborSearchTimeout = true
	configuration.Neighbors[0] = ""
	configuration.Neighbors[1] = ""
	configuration.Neighbors[2] = ""
	configuration.CommandEOL = 900
	configuration.SleepTimer = 10
	configuration.JitterValue = 2
	configuration.LastUpdate = time.Now()
	configuration.Identifier, e = uuid.NewRandom()

	if e != nil {
		panic(e)
	}
}

func deployInitialConfiguration() {
	neighborSearchTimeout = true
}
