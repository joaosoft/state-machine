# state-machine
[![Build Status](https://travis-ci.org/joaosoft/state-machine.svg?branch=master)](https://travis-ci.org/joaosoft/state-machine) | [![codecov](https://codecov.io/gh/joaosoft/state-machine/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/state-machine) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/state-machine)](https://goreportcard.com/report/github.com/joaosoft/state-machine) | [![GoDoc](https://godoc.org/github.com/joaosoft/state-machine?status.svg)](https://godoc.org/github.com/joaosoft/state-machine)

A simple state machine.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

>### Go
```
go get github.com/joaosoft/state-machine
```

## Usage 
This example is available in the project at [state-machine/examples](https://github.com/joaosoft/state-machine/tree/master/examples)

>### Configuration
#### State machine A (yaml)
```yaml
state_machine:
  -
    id: 1
    name: "New"
    transitions:
      -
        id: 2
        check:
          -
            "check_new_to_in-progress"
        execute:
          -
            "execute_new_to_in-progress"
  -
    id: 2
    name: "In progress"
    transitions:
      -
        id: 3
        check:
          -
            "check_in-progress_to_approved"
        execute:
          -
            "execute_in-progress_to_approved"
        events:
          success:
            -
              "event_success_in-progress_to_approved"
          error:
            -
              "event_error_in-progress_to_approved"
      -
        id: 4
        check:
          -
            "check_in-progress_to_denied"
        execute:
          -
            "execute_in-progress_to_denied"
        events:
          success:
            -
              "event_success_in-progress_to_denied"
          error:
            -
              "event_error_in-progress_to_denied"
  -
    id: 3
    name: "Approved"
  -
    id: 4
    name: "Denied"

users:
  operator:
    -
      id: 1
      transitions:
        -
          id: 2
          execute:
            - "execute_new_to_in-progress_user"
          events:
            success:
              - "event_success_new_to_in-progress_user"
    -
      id: 2
      transitions:
        -
          id: 3
        -
          id: 4
```

#### State machine B (json)
```json
{
  "state_machine": [
    {
      "id": 1,
      "name": "Todo",
      "transitions": [
        {
          "id": 2,
          "check": [
            "check_todo_to_in-development"
          ],
          "execute": [
            "execute_todo_to_in-development"
          ]
        }
      ]
    },
    {
      "id": 2,
      "name": "In development",
      "transitions": [
        {
          "id": 3,
          "check": [
            "check_in-development_to_done"
          ],
          "execute": [
            "execute_in-development_to_done"
          ],
          "events": {
            "success": [
              "event_success_in-development_to_done"
            ],
            "error": [
              "event_error_in-development_to_done"
            ]
          }
        },
        {
          "id": 4,
          "check": [
            "check_in-development_to_canceled"
          ],
          "execute": [
            "execute_in-development_to_canceled"
          ],
          "events": {
            "success": [
              "event_success_in-development_to_canceled"
            ],
            "error": [
              "event_error_in-development_to_canceled"
            ]
          }
        }
      ]
    },
    {
      "id": 3,
      "name": "Done"
    },
    {
      "id": 4,
      "name": "Canceled"
    }
  ],
  "users": {
    "worker": [
      {
        "id": 1,
        "transitions": [
          {
            "id": 2
          }
        ]
      },
      {
        "id": 2,
        "transitions": [
          {
            "id": 3
          },
          {
            "id": 4
          }
        ]
      }
    ]
  }
}
```

>### Implementation
```go
const (
	StateMachineA     state_machine.StateMachineType = "A"
	UserStateMachineA state_machine.UserType         = "operator"

	StateMachineB     state_machine.StateMachineType = "B"
	UserStateMachineB state_machine.UserType         = "worker"
)

func init() {
	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_new_to_in-progress", checkNewToInProgress, StateMachineA).
		AddCheckHandler("check_in-progress_to_approved", checkInProgressToApproved, StateMachineA).
		AddCheckHandler("check_in-progress_to_denied", checkInProgressToDenied, StateMachineA).
		//
		AddExecuteHandler("execute_new_to_in-progress", executeNewToInProgress, StateMachineA).
		AddExecuteHandler("execute_new_to_in-progress_user", executeNewToInProgressByUser, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_approved", executeInProgressToApproved, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_denied", executeInProgressToDenied, StateMachineA).
		//
		AddEventOnSuccessHandler("event_success_new_to_in-progress_user", eventOnSuccessNewToInProgressByUser, StateMachineA).
		AddEventOnSuccessHandler("event_success_new_to_in-progress", eventOnSuccessNewToInProgress, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_approved", eventOnSuccessInProgressToApproved, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_denied", eventOnSuccessInProgressToDenied, StateMachineA).
		//
		AddEventOnErrorHandler("event_error_new_to_in-progress", eventOnErrorNewToInProgress, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_approved", eventOnErrorInProgressToApproved, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_denied", eventOnErrorInProgressToDenied, StateMachineA).

		// state machine B
		AddManualHandler(state_machine.ManualInit, loadFromState, StateMachineB).
		AddCheckHandler("check_todo_to_in-development", checkTodoToInDevelopment, StateMachineB).
		AddCheckHandler("check_in-development_to_done", checkInDevelopmentToDone, StateMachineB).
		AddCheckHandler("check_in-development_to_canceled", checkInDevelopmentToCanceled, StateMachineB).
		//
		AddExecuteHandler("execute_todo_to_in-development", executeTodoToInDevelopment, StateMachineB).
		AddExecuteHandler("execute_in-development_to_canceled", executeInDevelopmentToCanceled, StateMachineB).
		AddExecuteHandler("execute_in-development_to_done", executeInDevelopmentToDone, StateMachineB).
		//
		AddEventOnSuccessHandler("event_success_todo_to_in-development", eventOnSuccessTodoToInDevelopment, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_done", eventOnSuccessInDevelopmentToDone, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_canceled", eventOnSuccessInDevelopmentToCanceled, StateMachineB).
		//
		AddEventOnErrorHandler("event_error_todo_to_in-development", eventOnErrorTodoToInDevelopment, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_done", eventOnErrorInDevelopmentToDone, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_canceled", eventOnErrorInDevelopmentToCanceled, StateMachineB)

	// add state machines
	// A
	if err := state_machine.NewStateMachine().
		Key(StateMachineA).
		File("/config/state_machines/state_machine_a.yaml").
		TransitionHandler(StateMachineATransitionHandler).
		Load(); err != nil {
		panic(err)
	}

	// B
	if err := state_machine.NewStateMachine().
		Key(StateMachineB).
		File("/config/state_machines/state_machine_b.json").
		TransitionHandler(StateMachineBTransitionHandler).
		Load(); err != nil {
		panic(err)
	}
}

func main() {
	stateMachines := []state_machine.StateMachineType{StateMachineA, StateMachineB}
	stateMachinesUsers := []state_machine.UserType{UserStateMachineA, UserStateMachineB}
	maxLen := 4
	ok := false

	// check transitions of state machines
	for index, stateMachine := range stateMachines {
		fmt.Printf("\n\n\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err := state_machine.NewCheckTransition().
					User(stateMachinesUsers[index]).
					StateMachine(stateMachine).
					From(i).
					To(j).
					Execute(1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\ntransition from %d to %d  with user %s ? %t", i, j, stateMachinesUsers[index], ok)
			}
		}
	}

	// get all transitions of state machine A
	transitions, err := state_machine.NewGetTransitions().
		User(UserStateMachineA).
		StateMachine(StateMachineA).
		From(1).
		Execute()
	if err != nil {
		panic(err)
	}
	for _, transition := range transitions {
		fmt.Printf("\ncan make transition to %s", transition.Name)
	}

	// execute transaction - state machine A
	ok, err = state_machine.NewTransition().
		User(UserStateMachineA).
		StateMachine(StateMachineA).
		From(1).
		To(2).
		Execute(1, "text", true)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}

	// execute transaction - state machine B
	ok, err = state_machine.NewTransition().
		User(UserStateMachineB).
		StateMachine(StateMachineB).
		To(2).
		Execute(1, "text", true)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}
}
```

>### Result
```
State Machine: A

transition from 1 to 4  with user operator ? false
transition from 1 to 3  with user operator ? false
check in-progress handler with [1 text true]
transition from 1 to 2  with user operator ? true
transition from 1 to 1  with user operator ? false
check in-progress to denied handler with [1 text true]
transition from 2 to 4  with user operator ? true
check in-progress to approved handler with [1 text true]
transition from 2 to 3  with user operator ? true
transition from 2 to 2  with user operator ? false
transition from 2 to 1  with user operator ? false
transition from 3 to 4  with user operator ? false
transition from 3 to 3  with user operator ? false
transition from 3 to 2  with user operator ? false
transition from 3 to 1  with user operator ? false
transition from 4 to 4  with user operator ? false
transition from 4 to 3  with user operator ? false
transition from 4 to 2  with user operator ? false
transition from 4 to 1  with user operator ? false


State Machine: B

load 'from' state handler with [1 text true]
transition from 1 to 4  with user worker ? false
load 'from' state handler with [1 text true]
transition from 1 to 3  with user worker ? false
load 'from' state handler with [1 text true]
check in-development handler with [1 text true]
transition from 1 to 2  with user worker ? true
load 'from' state handler with [1 text true]
transition from 1 to 1  with user worker ? false
load 'from' state handler with [1 text true]
transition from 2 to 4  with user worker ? false
load 'from' state handler with [1 text true]
transition from 2 to 3  with user worker ? false
load 'from' state handler with [1 text true]
check in-development handler with [1 text true]
transition from 2 to 2  with user worker ? true
load 'from' state handler with [1 text true]
transition from 2 to 1  with user worker ? false
load 'from' state handler with [1 text true]
transition from 3 to 4  with user worker ? false
load 'from' state handler with [1 text true]
transition from 3 to 3  with user worker ? false
load 'from' state handler with [1 text true]
check in-development handler with [1 text true]
transition from 3 to 2  with user worker ? true
load 'from' state handler with [1 text true]
transition from 3 to 1  with user worker ? false
load 'from' state handler with [1 text true]
transition from 4 to 4  with user worker ? false
load 'from' state handler with [1 text true]
transition from 4 to 3  with user worker ? false
load 'from' state handler with [1 text true]
check in-development handler with [1 text true]
transition from 4 to 2  with user worker ? true
load 'from' state handler with [1 text true]
transition from 4 to 1  with user worker ? false
can make transition to In progress
check in-progress handler with [1 text true]
execute in-progress handler with [1 text true]
by user: execute in-progress handler with [1 text true]
state machine: A, transition handler with [1 text true]
by user: success event in-progress handler with [1 text true]
load 'from' state handler with [1 text true]
check in-development handler with [1 text true]
execute in-development handler with [1 text true]
state machine: B, transition handler with [1 text true]
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
