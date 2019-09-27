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

func main() {
	var err error

	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_new_to_in-progress", CheckNewToInProgress, StateMachineA).
		AddCheckHandler("check_in-progress_to_approved", CheckInProgressToApproved, StateMachineA).
		AddCheckHandler("check_in-progress_to_denied", CheckInProgressToDenied, StateMachineA).
		//
		AddExecuteHandler("execute_new_to_in-progress", ExecuteNewToInProgress, StateMachineA).
		AddExecuteHandler("execute_new_to_in-progress_user", ExecuteNewToInProgressByUser, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_approved", ExecuteInProgressToApproved, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_denied", ExecuteInProgressToDenied, StateMachineA).
		//
		AddEventOnSuccessHandler("event_success_new_to_in-progress_user", EventOnSuccessNewToInProgressByUser, StateMachineA).
		AddEventOnSuccessHandler("event_success_new_to_in-progress", EventOnSuccessNewToInProgress, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_approved", EventOnSuccessInProgressToApproved, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_denied", EventOnSuccessInProgressToDenied, StateMachineA).
		//
		AddEventOnErrorHandler("event_error_new_to_in-progress", EventOnErrorNewToInProgress, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_approved", EventOnErrorInProgressToApproved, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_denied", EventOnErrorInProgressToDenied, StateMachineA).

		// state machine B
		AddCheckHandler("check_todo_to_in-development", CheckTodoToInDevelopment, StateMachineB).
		AddCheckHandler("check_in-development_to_done", CheckInDevelopmentToDone, StateMachineB).
		AddCheckHandler("check_in-development_to_canceled", CheckInDevelopmentToCanceled, StateMachineB).
		//
		AddExecuteHandler("execute_todo_to_in-development", ExecuteTodoToInDevelopment, StateMachineB).
		AddExecuteHandler("execute_in-development_to_canceled", ExecuteInDevelopmentToCanceled, StateMachineB).
		AddExecuteHandler("execute_in-development_to_done", ExecuteInDevelopmentToDone, StateMachineB).
		//
		AddEventOnSuccessHandler("event_success_todo_to_in-development", EventOnSuccessTodoToInDevelopment, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_done", EventOnSuccessInDevelopmentToDone, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_canceled", EventOnSuccessInDevelopmentToCanceled, StateMachineB).
		//
		AddEventOnErrorHandler("event_error_todo_to_in-development", EventOnErrorTodoToInDevelopment, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_done", EventOnErrorInDevelopmentToDone, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_canceled", EventOnErrorInDevelopmentToCanceled, StateMachineB)

	// add state machines
	// A
	if err = state_machine.NewStateMachine().
		Key(StateMachineA).
		File("/config/state_machines/state_machine_a.yaml").
		TransitionHandler(StateMachineATransitionHandler).
		Load(); err != nil {
		panic(err)
	}

	// B
	if err = state_machine.NewStateMachine().
		Key(StateMachineB).
		File("/config/state_machines/state_machine_b.json").
		TransitionHandler(StateMachineBTransitionHandler).
		Load(); err != nil {
		panic(err)
	}

	// check transitions of state machines
	stateMachines := []state_machine.StateMachineType{StateMachineA, StateMachineB}
	stateMachinesUsers := []state_machine.UserType{UserStateMachineA, UserStateMachineB}
	maxLen := 4
	ok := false

	for index, stateMachine := range stateMachines {
		fmt.Printf("\n\n\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err = state_machine.NewCheckTransition().
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

	// execute transaction
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

transition from 1 to 4  with user worker ? false
transition from 1 to 3  with user worker ? false
check in-development handler with [1 text true]
transition from 1 to 2  with user worker ? true
transition from 1 to 1  with user worker ? false
check in-development to canceled handler with [1 text true]
transition from 2 to 4  with user worker ? true
check in-development to done handler with [1 text true]
transition from 2 to 3  with user worker ? true
transition from 2 to 2  with user worker ? false
transition from 2 to 1  with user worker ? false
transition from 3 to 4  with user worker ? false
transition from 3 to 3  with user worker ? false
transition from 3 to 2  with user worker ? false
transition from 3 to 1  with user worker ? false
transition from 4 to 4  with user worker ? false
transition from 4 to 3  with user worker ? false
transition from 4 to 2  with user worker ? false
transition from 4 to 1  with user worker ? false
can make transition to In progress
check in-progress handler with [1 text true]
execute in-progress handler with [1 text true]
by user: execute in-progress handler with [1 text true]
state machine: A, transition handler with [1 text true]
by user: success event in-progress handler with [1 text true]
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
