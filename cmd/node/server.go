package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

//API Endpoints
// GET /api/pull: pull updated information from node
// POST /api/join: provide UUID value and push it throughout the network
// GET /api/reconnect: provide new neighbors

func pullConfig(w http.ResponseWriter, req *http.Request) {
	jsonConfig, err := json.Marshal(generateNewConfig())
	if err != nil {
		fmt.Print("Failed to JSONify config.")
	}
	w.Write(jsonConfig)
	//if currently timed out for checking for more neighbors add this server to neighbor list
	if neighborSearchTimeout {
		processNewNode(strings.Split(req.RemoteAddr, ":")[0])
	}
}

func joinNet(w http.ResponseWriter, req *http.Request) {
	data, _ := io.ReadAll(req.Body)
	UUIDString := string(data)
	newUUID, err := uuid.Parse(UUIDString)
	if err != nil {
		fmt.Println(err)
	}
	configuration.KnownNodes[UUIDString] = newUUID
}

func reconnect(w http.ResponseWriter, req *http.Request) {

}

func heartbeat(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "meowdy")
}
