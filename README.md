# state-machine
[![Build Status](https://travis-ci.org/joaosoft/state-machine.svg?branch=master)](https://travis-ci.org/joaosoft/state-machine) | [![codecov](https://codecov.io/gh/joaosoft/state-machine/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/state-machine) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/state-machine)](https://goreportcard.com/report/github.com/joaosoft/state-machine) | [![GoDoc](https://godoc.org/github.com/joaosoft/state-machine?status.svg)](https://godoc.org/github.com/joaosoft/state-machine)

A simple state machine.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* Users
* State Machines
* Transitions
* Handlers

## With configuration options
* WithLogger
* WithLogLevel

## With methods
* NewAddHandlers, to add handlers
* NewStateMachine, to add a new state machine
* NewCheckTransition, to check a transition
* NewTransition, to make a transition
* NewGetTransitions, to get the allowed transitions from the current state

## With manual handler types

### Check Transition (method NewCheckTransition)
* BeforeCheck, before check

### Execute Transition (method NewTransition)
* BeforeExecute, before execute
* AfterExecute, after execute
* OnSuccess, on success
* OnError, on error

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
<details>
    <summary>State machine A (yaml)</summary>
    ```yaml
    state_machine:
      -
        id: 1
        name: "New"
        transitions:
          -
            id: 2
            load:
              - "load_dummy"
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
</details>

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
	// :: add handlers

	// state machine A
	fmt.Println(":: State Machine: A - Adding handlers")
	state_machine.NewAddHandlers(StateMachineA).
		Load("load_dummy", loadDummy).
		//
		Check("check_new_to_in-progress", checkNewToInProgress).
		Check("check_in-progress_to_approved", checkInProgressToApproved).
		Check("check_in-progress_to_denied", checkInProgressToDenied).
		//
		Execute("execute_new_to_in-progress", executeNewToInProgress).
		Execute("execute_new_to_in-progress_user", executeNewToInProgressByUser).
		Execute("execute_in-progress_to_approved", executeInProgressToApproved).
		Execute("execute_in-progress_to_denied", executeInProgressToDenied).
		//
		EventSuccess("event_success_new_to_in-progress_user", eventOnSuccessNewToInProgressByUser).
		EventSuccess("event_success_new_to_in-progress", eventOnSuccessNewToInProgress).
		EventSuccess("event_success_in-progress_to_approved", eventOnSuccessInProgressToApproved).
		EventSuccess("event_success_in-progress_to_denied", eventOnSuccessInProgressToDenied).
		//
		EventError("event_error_new_to_in-progress", eventOnErrorNewToInProgress).
		EventError("event_error_in-progress_to_approved", eventOnErrorInProgressToApproved).
		EventError("event_error_in-progress_to_denied", eventOnErrorInProgressToDenied)

	// state machine B
	fmt.Println(":: State Machine: B - Adding handlers")
	state_machine.NewAddHandlers(StateMachineB).
		Manual(beforeExecuteLoadFromState, state_machine.BeforeCheck, state_machine.BeforeExecute).
		//
		Check("check_todo_to_in-development", checkTodoToInDevelopment).
		Check("check_in-development_to_done", checkInDevelopmentToDone).
		Check("check_in-development_to_canceled", checkInDevelopmentToCanceled).
		//
		Execute("execute_todo_to_in-development", executeTodoToInDevelopment).
		Execute("execute_in-development_to_canceled", executeInDevelopmentToCanceled).
		Execute("execute_in-development_to_done", executeInDevelopmentToDone).
		//
		EventSuccess("event_success_todo_to_in-development", eventOnSuccessTodoToInDevelopment).
		EventSuccess("event_success_in-development_to_done", eventOnSuccessInDevelopmentToDone).
		EventSuccess("event_success_in-development_to_canceled", eventOnSuccessInDevelopmentToCanceled).
		//
		EventError("event_error_todo_to_in-development", eventOnErrorTodoToInDevelopment).
		EventError("event_error_in-development_to_done", eventOnErrorInDevelopmentToDone).
		EventError("event_error_in-development_to_canceled", eventOnErrorInDevelopmentToCanceled)

	// :: add state machines

	// A
	fmt.Println(":: State Machine: A - Adding state machine")
	if err := state_machine.NewStateMachine().
		Key(StateMachineA).
		File("/config/state_machines/state_machine_a.yaml").
		TransitionHandler(StateMachineATransitionHandler).
		Load(); err != nil {
		panic(err)
	}

	// B
	fmt.Println(":: State Machine: B - Adding state machine")
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

	// get all transitions of state machine A
	fmt.Println("\n:: State Machine: A - get all transition from 1 to 2")
	transitions, err := state_machine.NewGetTransitions().
		User(UserStateMachineA).
		StateMachine(StateMachineA).
		From(1).
		Execute()
	if err != nil {
		panic(err)
	}
	for _, transition := range transitions {
		fmt.Printf("can make transition to %s\n", transition.Name)
	}

	// check transitions of state machines
	for index, stateMachine := range stateMachines {
		fmt.Printf("\n:: State Machine: %s - check transitions\n", stateMachine)
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
				fmt.Printf("transition from %d to %d  with user %s ? %t\n", i, j, stateMachinesUsers[index], ok)
			}
		}
	}

	// check transaction - state machine B - from the state loaded by method 'beforeExecuteLoadFromState' to state 2
	fmt.Println("\n:: State Machine: B - check transition from state 1 (loaded) to state 2")
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

	// execute transaction - state machine B - from the state loaded by method 'beforeExecuteLoadFromState' to state 2
	fmt.Println("\n:: State Machine: B - making transition from state 1 (loaded) to state 2")
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
:: State Machine: A - Adding handlers
:: State Machine: B - Adding handlers
:: State Machine: A - Adding state machine
:: State Machine: B - Adding state machine

:: State Machine: A - get all transition from 1 to 2
can make transition to In progress

:: State Machine: A - check transitions
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

:: State Machine: B - check transitions
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

:: State Machine: B - check transition from state 1 (loaded) to state 2
load 'from' state handler with [1 text true]
check in-development handler with [1 text true]
execute in-development handler with [1 text true]
state machine: B, transition handler with [1 text true]

:: State Machine: B - making transition from state 1 (loaded) to state 2
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

## License
Released under the [MIT License](http://opensource.org/licenses/MIT).