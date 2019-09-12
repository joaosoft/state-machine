package state_machine

import (
	"os"
)

func GetEnv() string {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	return env
}
