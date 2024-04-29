package config

import "github.com/google/uuid"

type config struct {
	identifier uuid
}

func generateConfig() *config {
	conf := config{
		identifier: uuid.NewRandom(),
	}

	return &conf
}
