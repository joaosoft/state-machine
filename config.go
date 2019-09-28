package state_machine

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// appConfig ...
type appConfig struct {
	StateMachine *stateMachineConfig `json:"state_machine"`
}

// stateMachineConfig ...
type stateMachineConfig struct {
	Log       struct {
		Level string `json:"level"`
	} `json:"log"`
}

// newConfig ...
func newConfig() (*appConfig, manager.IConfig, error) {
	appConfig := &appConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig)

	return appConfig, simpleConfig, err
}
