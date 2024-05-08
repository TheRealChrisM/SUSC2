package main

import (
	"os"
	"os/exec"
	"strings"
)

func run(cmd Command) CommandOutput {
	output := CommandOutput{
		Id:     cmd.Identifier,
		Stdout: "",
		Stderr: "",
	}
	args := strings.Fields(os.ExpandEnv(cmd.Command))
	toRun := exec.Command(args[0], args[1:len(args)]...)
	var stdout []byte
	var stderr []byte
	err := toRun.Run()
	if err != nil {
		output.Stdout = ""
		output.Stderr = err.Error()
		return output
	} else {
		_, err2 := toRun.Stdout.Write(stdout)
		_, err3 := toRun.Stderr.Write(stderr)
		if err2 != nil {
			output.Stdout = ""
			output.Stderr = err2.Error()
		} else if err3 != nil {
			output.Stdout = ""
			output.Stderr = err3.Error()
		} else {
			output.Stdout = string(stdout[:])
			output.Stderr = string(stderr[:])
		}
	}
	return output
}
