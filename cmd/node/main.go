package main

import (
	"flag"
	"fmt"
	"net/http"
)

var server string

func main() {
	flag.StringVar(&server, "server", "", "(optional) address of a node whose network you want to join (e.g. \"http://192.168.1.1:8443\")")
	flag.Parse()
	fmt.Print(server)
	if confirmServerValueProvided() {
		if validateServerAddress(server) {
			deployInitialConfiguration()
		}
	} else {
		bootstrapSelf()
	}
	http.HandleFunc("/api/pull", pullConfig)
	http.HandleFunc("/api/join", joinNet)
	http.HandleFunc("/api/reconnect", reconnect)
	http.HandleFunc("/api", heartbeat)
	http.ListenAndServe("0.0.0.0:31337", nil)
	fmt.Print(configuration.LastUpdate)
}
