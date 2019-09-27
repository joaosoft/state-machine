package state_machine

import (
	"sync"

	"github.com/joaosoft/logger"
)

type StateMap map[int]*State

type Context struct {
	User         UserType
	StateMachine StateMachineType
	From         int
	To           int
	resource     interface{}
}

type CheckHandler func(ctx *Context, args ...interface{}) (bool, error)
type ExecuteHandler func(ctx *Context, args ...interface{}) error
type EventSuccessHandler func(ctx *Context, args ...interface{})
type EventErrorHandler func(ctx *Context, err error, args ...interface{})
type TransitionHandler func(ctx *Context, args ...interface{}) error

type CheckHandlerMap map[string]CheckHandler
type ExecuteHandlerMap map[string]ExecuteHandler
type EventSuccessHandlerMap map[string]EventSuccessHandler
type EventErrorHandlerMap map[string]EventErrorHandler
type StateMachineHandlersMap map[StateMachineType]*HandlersMap

type Handlers struct {
	handlersMap             *HandlersMap
	stateMachineHandlersMap StateMachineHandlersMap
}

type stateMachine struct {
	config              *StateMachineConfig
	stateMachineMap     StateMachineMap
	userStateMachineMap UserStateMachineMap
	handlers            *Handlers
	logger              logger.ILogger
	mux                 *sync.RWMutex
}

type StateMachineMap map[StateMachineType]*StateMachineData
type StateMachineData struct {
	stateMap          StateMap
	transitionHandler TransitionHandler
}
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
	Check   CheckHandlerList   `json:"check"`
	Execute ExecuteHandlerList `json:"execute"`
	Events  Events             `json:"events"`
}

type Events struct {
	Success EventSuccessHandlerList `json:"success"`
	Error   EventErrorHandlerList   `json:"error"`
}

type CheckHandlerList []CheckHandler
type ExecuteHandlerList []ExecuteHandler
type EventSuccessHandlerList []EventSuccessHandler
type EventErrorHandlerList []EventErrorHandler

type HandlersMap struct {
	Check   map[string]CheckHandler   `json:"check"`
	Execute map[string]ExecuteHandler `json:"execute"`
	Events  *EventMap                 `json:"events"`
}

type EventMap struct {
	Success map[string]EventSuccessHandler `json:"success"`
	Error   map[string]EventErrorHandler   `json:"error"`
}

type UserType string
type StateMachineType string
