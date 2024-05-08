package main

import (
	"time"

	"github.com/google/uuid"
)

type Command struct {
	Target     uuid.UUID `json:"target"`
	Command    string    `json:"command"`
	DeployTime time.Time `json:"deploy_time"`
	Identifier uuid.UUID `json:"identifier"`
}

type CommandOutput struct {
	Id     uuid.UUID `json:"id"`
	Stdout string    `json:"stdout"`
	Stderr string    `json:"stderr"`
}

type Config struct {
	Identifier  uuid.UUID            `json:"identifier"`
	Neighbors   [3]string            `json:"neighbors"`
	KnownNodes  map[string]uuid.UUID `json:"KnownNodes"`
	TaskList    map[string]Command   `json:"task_list"`
	CommandEOL  int                  `json:"command_end_of_life"`
	SleepTimer  int                  `json:"sleep_timer"`
	JitterValue int                  `json:"jitter_value"`
	LastUpdate  time.Time            `json:"last_update"`
}

type Request struct {
	Id            uuid.UUID       `json:"id"`
	LastConnected map[string]int  `json:"last_connected"`
	CommandOutput []CommandOutput `json:"command_output"`
}

type Response struct {
	Config   Config    `json:"config"`
	Commands []Command `json:"commands"`
}
