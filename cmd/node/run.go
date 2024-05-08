package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func run(cmd Command) {
	fmt.Println(cmd)
	args := strings.Fields(os.ExpandEnv(cmd.Command))
	toRun := exec.Command(args[0], args[1:len(args)]...)
	toRun.Run()
}

func runEverything() {
	for {
		updateInformation()
		for k, v := range configuration.TaskList {
			_, ok := completedTasks[k]
			if ok {
				delete(configuration.TaskList, k)
			} else {
				fmt.Println(v.Target.String())
				fmt.Println(configuration.Identifier.String())
				fmt.Println(v.Target.String() == configuration.Identifier.String())
				if v.Target.String() == configuration.Identifier.String() {
					fmt.Printf("Running following command: %s", v)
					run(v)
					completedTasks[k] = v
					delete(configuration.TaskList, k)
				}
			}
		}
		fmt.Print("Sleeping, current configuration is: ")
		fmt.Println(configuration)
		time.Sleep(time.Duration(configuration.SleepTimer) * time.Second)
	}
}
