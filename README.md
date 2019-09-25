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
#### State machine A
```json
{
  "state_machine": [
    {
      "id": 1,
      "name": "New",
      "transitions": [
        {
          "id": 2,
          "check": [
            "check_in-progress"
          ],
          "execute": [
            "execute_in-progress"
          ]
        }
      ]
    },
    {
      "id": 2,
      "name": "In progress",
      "transitions": [
        {
          "id": 3,
          "check": [
            "check_in-progress_to_approved"
          ],
          "execute": [
            "execute_approved"
          ],
          "events": {
            "success": [
              "event_success_approved"
            ],
            "error": [
              "event_error_approved"
            ]
          }
        },
        {
          "id": 4,
          "check": [
            "check_in-progress_to_denied"
          ],
          "execute": [
            "execute_denied"
          ],
          "events": {
            "success": [
              "event_success_denied"
            ],
            "error": [
              "event_error_denied"
            ]
          }
        }
      ]
    },
    {
      "id": 3,
      "name": "Approved"
    },
    {
      "id": 4,
      "name": "Denied"
    }
  ],
  "users": {
    "operator": [
      {
        "id": 1,
        "transitions": [2]
      },
      {
        "id": 2,
        "transitions": [3, 4]
      }
    ]
  }
}
```

#### State machine B
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
            "check_in-development"
          ],
          "execute": [
            "execute_in-development"
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
            "execute_in-development"
          ],
          "events": {
            "success": [
              "event_success_in-development"
            ],
            "error": [
              "event_error_in-development"
            ]
          }
        },
        {
          "id": 4,
          "check": [
            "check_in-development_to_canceled"
          ],
          "execute": [
            "execute_canceled"
          ],
          "events": {
            "success": [
              "event_success_canceled"
            ],
            "error": [
              "event_error_canceled"
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
        "transitions": [2]
      },
      {
        "id": 2,
        "transitions": [3, 4]
      }
    ]
  }
}
```

>### Implementation
```go
const (
	StateMachineA     = "A"
	UserStateMachineA = "operator"

	StateMachineB     = "B"
	UserStateMachineB = "worker"
)

func main() {
	var err error

	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_in-progress", CheckInProgress).
		AddExecuteHandler("execute_in-progress", ExecuteInProgress).
		AddEventOnSuccessHandler("event_success_in-progress", EventOnSuccessInProgress).
		AddEventOnErrorHandler("event_error_in-progress", EventOnErrorInProgress).

		// state machine B
		AddCheckHandler("check_in-development", CheckInDevelopment).
		AddExecuteHandler("execute_in-development", ExecuteInDevelopment).
		AddEventOnSuccessHandler("event_success_in-development", EventOnSuccessInDevelopment).
		AddEventOnErrorHandler("event_error_in-development", EventOnErrorInDevelopment)

	// add state machines
	if err = state_machine.AddStateMachine(StateMachineA, "/config/state_machine_a.json"); err != nil {
		panic(err)
	}
	if err = state_machine.AddStateMachine(StateMachineB, "/config/state_machine_b.json"); err != nil {
		panic(err)
	}

	// check transitions of state machines
	stateMachines := []string{StateMachineA, StateMachineB}
	stateMachinesUsers := []string{UserStateMachineA, UserStateMachineB}
	maxLen := 5
	ok := false

	for index, stateMachine := range stateMachines {
		fmt.Printf("\n\n\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err = state_machine.CheckTransition(stateMachine, stateMachinesUsers[index], i, j, 1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\ntransition from %d to %d  with user %s ? %t", i, j, stateMachinesUsers[index], ok)
			}
		}
	}

	// get all transitions of state machine A
	transitions, err := state_machine.GetTransitions(StateMachineA, UserStateMachineA, 1)
	if err != nil {
		panic(err)
	}
	for _, transition := range transitions {
		fmt.Printf("\ncan make transition to %s", transition.Name)
	}

	// execute transaction
	ok, err = state_machine.ExecuteTransition(StateMachineA, UserStateMachineA, 1, 2)
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

transition from 1 to 5  with user operator ? false
transition from 1 to 4  with user operator ? false
transition from 1 to 3  with user operator ? false
check in-progress handler with [[1 text true]]
transition from 1 to 2  with user operator ? false
transition from 1 to 1  with user operator ? false
transition from 2 to 5  with user operator ? false
transition from 2 to 4  with user operator ? false
transition from 2 to 3  with user operator ? false
transition from 2 to 2  with user operator ? false
transition from 2 to 1  with user operator ? false
transition from 3 to 5  with user operator ? false
transition from 3 to 4  with user operator ? false
transition from 3 to 3  with user operator ? false
transition from 3 to 2  with user operator ? false
transition from 3 to 1  with user operator ? false
transition from 4 to 5  with user operator ? false
transition from 4 to 4  with user operator ? false
transition from 4 to 3  with user operator ? false
transition from 4 to 2  with user operator ? false
transition from 4 to 1  with user operator ? false
transition from 5 to 5  with user operator ? false
transition from 5 to 4  with user operator ? false
transition from 5 to 3  with user operator ? false
transition from 5 to 2  with user operator ? false
transition from 5 to 1  with user operator ? false


State Machine: B

transition from 1 to 5  with user worker ? false
transition from 1 to 4  with user worker ? false
transition from 1 to 3  with user worker ? false
check in-development handler with [[1 text true]]
transition from 1 to 2  with user worker ? false
transition from 1 to 1  with user worker ? false
transition from 2 to 5  with user worker ? false
transition from 2 to 4  with user worker ? false
transition from 2 to 3  with user worker ? false
transition from 2 to 2  with user worker ? false
transition from 2 to 1  with user worker ? false
transition from 3 to 5  with user worker ? false
transition from 3 to 4  with user worker ? false
transition from 3 to 3  with user worker ? false
transition from 3 to 2  with user worker ? false
transition from 3 to 1  with user worker ? false
transition from 4 to 5  with user worker ? false
transition from 4 to 4  with user worker ? false
transition from 4 to 3  with user worker ? false
transition from 4 to 2  with user worker ? false
transition from 4 to 1  with user worker ? false
transition from 5 to 5  with user worker ? false
transition from 5 to 4  with user worker ? false
transition from 5 to 3  with user worker ? false
transition from 5 to 2  with user worker ? false
transition from 5 to 1  with user worker ? false
can make transition to In progress
check in-progress handler with [[]]transition !ok
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
