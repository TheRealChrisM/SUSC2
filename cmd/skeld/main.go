package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TheRealChrisM/SUSC2/pkg/skserver"
)

func main() {

	cmds := make(chan string)
	cmds <- "echo hello"
	fmt.Print(cmds)
	skserver.Serve(&cmds)

	var s = bufio.NewScanner(os.Stdin)

	for s.Scan() {
		cmds <- s.Text()
		fmt.Print(cmds)
	}
}
