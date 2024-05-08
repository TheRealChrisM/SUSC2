package main

import (
	"flag"
	"net/http"
)

var server string

func main() {
	flag.StringVar(&server, "server", "", "(optional) address of a node whose network you want to join (e.g. \"192.168.1.1\")")
	flag.Parse()
	if confirmServerValueProvided() {
		//if validateServerAddress(server) {
		deployInitialConfiguration()
		//}
	} else {
		bootstrapSelf()
	}
	http.HandleFunc("/api/pull", pullConfig)
	http.HandleFunc("/api/join", joinNet)
	http.HandleFunc("/api/reconnect", reconnect)
	//http.HandleFunc("/api/ctrl", issueCommand)
	http.HandleFunc("/api", heartbeat)
	http.ListenAndServe("0.0.0.0:31337", nil)
}
