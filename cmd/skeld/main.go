package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// type commandResponse struct {
// 	err string
// 	cid int
// 	cmd string
// }

var commands []string

func offer_command(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cid, err := strconv.ParseUint(r.Form.Get("cid"), 10, 32)
	var res interop
	if err != nil {
		res.cid = -1
		res.err = "Invalid cid"
	} else if int(cid) >= len(commands) || int(cid) < 0 {
		res.cid = len(commands)
		res.err = "Not yet supported"
	} else {
		res.cid = int(cid)
		res.cmd = commands[cid]
	}
	bs, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	w.Write(bs)
}

func main() {

	for {

	}
}
