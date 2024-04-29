package interop

import (
	"github.com/google/uuid"
)

type Config struct {
	Identifier   uuid.UUID
	Interval     int
	KnownServers []string
}
type CommandResponse struct {
	Err string
	Cid int
	Cmd string
}

func GenerateConfig(myip string) Config {
	var generatedUUID, e = uuid.NewRandom()

	if e != nil {
		panic(e)
	}

	return Config{
		Identifier:   generatedUUID,
		Interval:     10,
		KnownServers: []string{myip + ":8443"},
	}
}
