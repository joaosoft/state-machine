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
	Resource     interface{}
	Args         []interface{}
}

type ManualHandler func(ctx *Context) error

type LoadHandler func(ctx *Context) error
type CheckHandler func(ctx *Context) (bool, error)
type ExecuteHandler func(ctx *Context) error
type EventSuccessHandler func(ctx *Context)
type EventErrorHandler func(ctx *Context, err error)
type TransitionHandler func(ctx *Context) error

type ManualHandlerMap map[ManualHandlerTag]ManualHandlerList
type ManualHandlerList []ManualHandler

type LoadHandlerMap map[string]LoadHandler
type CheckHandlerMap map[string]CheckHandler
type ExecuteHandlerMap map[string]ExecuteHandler
type EventSuccessHandlerMap map[string]EventSuccessHandler
type EventErrorHandlerMap map[string]EventErrorHandler
type StateMachineHandlersMap map[StateMachineType]*HandlersMap

type handlers struct {
	handlersMap             *HandlersMap
	stateMachineHandlersMap StateMachineHandlersMap
}

type stateMachine struct {
	config              *StateMachineConfig
	stateMachineMap     StateMachineMap
	userStateMachineMap UserStateMachineMap
	handlers            *handlers
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
	Id            int
	Name          string
	TransitionMap TransitionMap
}

type TransitionMap map[int]*Transition

type Transition struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Handler Handler `json:"-"`
}

type Handler struct {
	Load    LoadHandlerList
	Check   CheckHandlerList
	Execute ExecuteHandlerList
	Events  Events
}

type Events struct {
	Success EventSuccessHandlerList
	Error   EventErrorHandlerList
}

type LoadHandlerList []LoadHandler
type CheckHandlerList []CheckHandler
type ExecuteHandlerList []ExecuteHandler
type EventSuccessHandlerList []EventSuccessHandler
type EventErrorHandlerList []EventErrorHandler

type HandlersMap struct {
	Manual  ManualHandlerMap
	Load    LoadHandlerMap
	Check   CheckHandlerMap
	Execute ExecuteHandlerMap
	Events  *EventMap
}

type EventMap struct {
	Success EventSuccessHandlerMap
	Error   EventErrorHandlerMap
}

type UserType string
type StateMachineType string
