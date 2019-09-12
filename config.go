package state_machine

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	StateMachine *StateMachineConfig `json:"state_machine"`
}

// StateMachineConfig ...
type StateMachineConfig struct {
	States    []StateConfig
	Log       struct {
		Level string `json:"level"`
	} `json:"log"`
}

type StateConfig struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Transitions []TransitionConfig `json:"transitions"`
}

type TransitionConfig struct {
	Id      int    `json:"id"`
	Handler string `json:"handler"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
