package state_machine

import (
	"sync"

	"github.com/joaosoft/logger"
)

type StateMap map[int]*State

type CheckHandler func(args ...interface{}) (bool, error)
type ExecuteHandler func(args ...interface{}) (bool, error)
type EventHandler func(args ...interface{}) error

type CheckHandlerMap map[string]CheckHandler
type ExecuteHandlerMap map[string]ExecuteHandler
type EventHandlerMap map[string]EventHandler
type StateMachineHandlersMap map[StateMachineType]*HandlersMap

type Handlers struct {
	handlersMap             *HandlersMap
	stateMachineHandlersMap StateMachineHandlersMap
}

type StateMachine struct {
	config              *StateMachineConfig
	userStateMachineMap UserStateMachineMap
	handlers            *Handlers
	logger              logger.ILogger
	mux                 *sync.RWMutex
}

type StateMachineMap map[StateMachineType]StateMap
type UserStateMachineMap map[UserType]StateMachineMap

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

type HandlersMap struct {
	Check   map[string]CheckHandler   `json:"check"`
	Execute map[string]ExecuteHandler `json:"execute"`
	Events  *EventMap                 `json:"events"`
}

type EventMap struct {
	Success map[string]EventHandler `json:"success"`
	Error   map[string]EventHandler `json:"error"`
}

type UserType string
type StateMachineType string
