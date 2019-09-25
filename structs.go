package state_machine

import (
	"github.com/joaosoft/logger"
)

type StateMachineCfg struct {
	StateMachine []struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Transitions []struct {
			Id      int      `json:"id"`
			Check   []string `json:"check"`
			Execute []string `json:"execute"`
			Events  struct {
				Success []string `json:"success"`
				Error   []string `json:"error"`
			} `json:"events"`
		} `json:"transitions"`
	} `json:"state_machine"`
	Users map[string][]struct {
		Id          int   `json:"id"`
		Transitions []int `json:"transitions"`
	} `json:"users"`
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
}

type StateMachineMap map[string]StateMap
type UserStateMachine map[string]StateMachineMap

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
