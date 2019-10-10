package state_machine

import (
	"sync"

	"github.com/joaosoft/logger"
)

type stateMap map[string]*state

type Context struct {
	Role         RoleType
	StateMachine StateMachineType
	From         string
	To           string
	Resource     interface{}
	Args         []interface{}
}

type ManualHandler func(ctx *Context) error

type LoadHandler func(ctx *Context) error
type CheckHandler func(ctx *Context) (bool, error)
type ExecuteHandler func(ctx *Context) error
type EventSuccessHandler func(ctx *Context) error
type EventErrorHandler func(ctx *Context, err error) error
type TransitionHandler func(ctx *Context) error

type manualHandlerMap map[manualHandlerKey]manualHandlerList
type manualHandlerList []ManualHandler

type loadHandlerMap map[string]LoadHandler
type checkHandlerMap map[string]CheckHandler
type executeHandlerMap map[string]ExecuteHandler
type eventSuccessHandlerMap map[string]EventSuccessHandler
type eventErrorHandlerMap map[string]EventErrorHandler
type stateMachineHandlersMap map[StateMachineType]*handlersMap

type handlers struct {
	handlersMap             *handlersMap
	stateMachineHandlersMap stateMachineHandlersMap
}

type stateMachine struct {
	config              *stateMachineConfig
	stateMachineMap     stateMachineMap
	roleStateMachineMap roleStateMachineMap
	handlers            *handlers
	logger              logger.ILogger
	mux                 *sync.RWMutex
}

type stateMachineMap map[StateMachineType]*stateMachineData
type stateMachineData struct {
	stateMap          stateMap
	transitionHandler TransitionHandler
}
type roleStateMachineMap map[RoleType]stateMachineMap

type state struct {
	id            string
	name          string
	transitionMap transitionMap
}

type transitionMap map[string]*Transition

type Transition struct {
	Id      string    `json:"id"`
	Name    string `json:"name"`
	handler handler
}

type handler struct {
	load    loadHandlerList
	check   checkHandlerList
	execute executeHandlerList
	events  events
}

type events struct {
	success eventSuccessHandlerList
	error   eventErrorHandlerList
}

type loadHandlerList []LoadHandler
type checkHandlerList []CheckHandler
type executeHandlerList []ExecuteHandler
type eventSuccessHandlerList []EventSuccessHandler
type eventErrorHandlerList []EventErrorHandler

type handlersMap struct {
	manual  manualHandlerMap
	Load    loadHandlerMap
	check   checkHandlerMap
	execute executeHandlerMap
	events  *eventMap
}

type eventMap struct {
	success eventSuccessHandlerMap
	error   eventErrorHandlerMap
}

type RoleType string
type StateMachineType string
