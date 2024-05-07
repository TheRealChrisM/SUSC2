package node

import "github.com/google/uuid"

type Command struct {
	Id      uuid.UUID `json:"id"`
	Command string    `json:"command"`
	Parent  string    `json:"parent"`
	Target  string    `json:"target"`
}

type CommandOutput struct {
	Id     uuid.UUID `json:"id"`
	Stdout string    `json:"stdout"`
	Stderr string    `json:"stderr"`
}

type Config struct {
	Id       uuid.UUID `json:"id"`
	Servers  []string  `json:"servers"`
	Interval int       `json:"interval"`
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
