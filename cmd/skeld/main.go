package main

import (
	"bufio"
	"os"

	"github.com/TheRealChrisM/SUSC2/pkg/skserver"
)

func main() {

	cmds := make(chan string)
	skserver.Serve(&cmds)

	var s = bufio.NewScanner(os.Stdin)

	for s.Scan() {
		cmds <- s.Text()
	}
}
