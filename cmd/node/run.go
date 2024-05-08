package main

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

func run(cmd Command) {
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
}

func runEverything() {
	for {
		updateInformation()
		for k, v := range configuration.TaskList {
			_, ok := completedTasks[k]
			if ok {
				delete(configuration.TaskList, k)
			} else {
				if v.Target == configuration.Identifier {
					run(v)
					completedTasks[k] = v
					delete(configuration.TaskList, k)
				}
			}
		}
		time.Sleep(time.Duration(configuration.SleepTimer) * time.Second)
	}
}
