package state_machine

import (
	"sync"

	"github.com/joaosoft/logger"
)

type StateMachineCfg struct {
	StateMachine []struct {
		Id          int    `json:"id" yaml:"id"`
		Name        string `json:"name" yaml:"name"`
		Transitions []struct {
			Id      int      `json:"id" yaml:"id"`
			Check   []string `json:"check" yaml:"check"`
			Execute []string `json:"execute" yaml:"execute"`
			Events  struct {
				Success []string `json:"success" yaml:"success"`
				Error   []string `json:"error" yaml:"error"`
			} `json:"events" yaml:"events"`
		} `json:"transitions" yaml:"transitions"`
	} `json:"state_machine" yaml:"state_machine"`
	Users map[string][]struct {
		Id          int   `json:"id" yaml:"id"`
		Transitions []int `json:"transitions" yaml:"transitions"`
	} `json:"users" yaml:"users"`
}

type StateMap map[int]*State

type CheckHandler func(args ...interface{}) (bool, error)
type ExecuteHandler func(args ...interface{}) (bool, error)
type EventHandler func(args ...interface{}) error

type StateMachine struct {
	config              *StateMachineConfig
	stateMachineMap     StateMachineMap
	userStateMachineMap UserStateMachine
	handlerMap          HandlerMap
	logger              logger.ILogger
	mux                 *sync.RWMutex
}

type StateMachineMap map[StateMachineType]StateMap
type UserStateMachine map[UserType]StateMachineMap

type State struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	TransitionMap TransitionMap `json:"transitions"`
}

type TransitionMap map[int]*Transition

type Transition struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Handler Handler `json:"handler"`
}

type Handler struct {
	Check   []CheckHandler   `json:"check"`
	Execute []ExecuteHandler `json:"execute"`
	Events  Events           `json:"events"`
}

type Events struct {
	Success []EventHandler `json:"success"`
	Error   []EventHandler `json:"error"`
}

type HandlerMap struct {
	Check   map[string]CheckHandler   `json:"check"`
	Execute map[string]ExecuteHandler `json:"execute"`
	Events  EventMap                  `json:"events"`
}

type EventMap struct {
	Success map[string]EventHandler `json:"success"`
	Error   map[string]EventHandler `json:"error"`
}

type UserType string
type StateMachineType string
