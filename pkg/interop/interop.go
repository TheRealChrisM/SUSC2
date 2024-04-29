package interop

import (
	"errors"

	"github.com/google/uuid"
)

type config struct {
	identifier uuid.UUID
	beaconTime int
}
type commandResponse struct {
	err string
	cid int
	cmd string
}

func generateConfig() *config {
	var generatedUUID, e = uuid.NewRandom()

	if e != nil {
		errors.New("Failed to generate UUID.")
	}

	conf := config{
		identifier: generatedUUID,
		beaconTime: 10,
	}

	return &conf
}

// func main() {
// 	var meow = generateConfig()

// 	fmt.Print(meow)
// }
