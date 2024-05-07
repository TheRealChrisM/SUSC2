package node

import (
	"errors"
	"fmt"
	"regexp"
)

var neighbors map[string]int

// var clients
var interval int

func validateServerAddress(address string) {
	match, err := regexp.MatchString("^http://[0-9]\\.[0-9]\\.[0-9]\\.[0-9]:[0-9]+$", address)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Invalid server address: %s", address)))
	}
	if !match {
		panic(errors.New(fmt.Sprintf("Invalid server address: %s", address)))
	}
}

func isUnreachable(address string) bool {
	neighbors
}

func bootstrapSelf() {

}
