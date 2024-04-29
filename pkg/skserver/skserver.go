package skserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/TheRealChrisM/SUSC2/pkg/interop"
)

func getCommand(server string, cid int) (string, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	hresp, err := client.PostForm(server+"/pull", url.Values{
		"cid": {strconv.Itoa(cid)},
	})
	if err != nil {
		return "", err
	}
	resp, err := io.ReadAll(hresp.Body)
	if err != nil {
		return "", err
	}
	var res interop.CommandResponse
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return "", err
	}
	if len(res.Err) > 0 { // TODO upgrade bidirectional
		return "", errors.New(res.Err)
	}
	return res.Cmd, nil
}

func Fetch(config *interop.Config, configMutex *sync.RWMutex, cmds *chan string) {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	i := 0
	for {
		for {
			configMutex.RLock()
			srv := config.KnownServers[rng.Intn(len(config.KnownServers))]
			configMutex.RUnlock()

			cmd, err := getCommand(srv, i)
			if err != nil {
				continue
			}
			*cmds <- cmd
			break
		}
		i++
	}
}

var commands []string
var mu sync.RWMutex

func offerCommand(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	cid, err := strconv.ParseUint(r.Form.Get("cid"), 10, 32)
	var res interop.CommandResponse
	mu.RLock()
	if err != nil {
		res.Cid = -1
		res.Err = "Invalid cid"
	} else if int(cid) >= len(commands) || int(cid) < 0 {
		res.Cid = len(commands) - 1
		res.Err = "Not yet"
	} else {
		res.Cid = int(cid)
		res.Cmd = commands[cid]
	}
	mu.RUnlock()
	bs, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	w.Write(bs)
}

func offerConfig(w http.ResponseWriter, r *http.Request) {
	bs, err := json.Marshal(interop.GenerateConfig("0.0.0.0")) // TODO get current IP
	if err != nil {
		fmt.Printf("Failed!")
	} else {
		w.Write(bs)
	}
}

func Serve(cmds *chan string) error {
	http.HandleFunc("/setup", offerConfig)
	http.HandleFunc("/pull", offerCommand)

	go func() {
		for s := range *cmds {
			mu.Lock()
			commands = append(commands, s)
			mu.Unlock()
		}
	}()
	return http.ListenAndServe("0.0.0.0:8443", nil)
}
