# state-machine
[![Build Status](https://travis-ci.org/joaosoft/state-machine.svg?branch=master)](https://travis-ci.org/joaosoft/state-machine) | [![codecov](https://codecov.io/gh/joaosoft/state-machine/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/state-machine) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/state-machine)](https://goreportcard.com/report/github.com/joaosoft/state-machine) | [![GoDoc](https://godoc.org/github.com/joaosoft/state-machine?status.svg)](https://godoc.org/github.com/joaosoft/state-machine)

A simple state machine checker.

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
```json
{
  "state_machine": {
    "log": {
      "level": "info"
    },
    "states": [
      {
        "id": 1,
        "name": "New",
        "transitions": [
          {
            "id":  2
          }
        ]
      },
      {
        "id": 2,
        "name": "In progress",
        "transitions": [
          {
            "id":  3,
            "handler": "check_3"
          },
          {
            "id":  4
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
    ]
  },
  "manager": {
    "log": {
      "level": "info"
    }
  }
}
```

>### Implementation
```go
const (
	StateMachineA = "A"
	StateMachineB = "B"
)

func main() {
	var err error

	// add transition check handlers
	state_machine.
		AddTransitionCheckHandler("check_in-progress", CheckInProgress).
		AddTransitionCheckHandler("check_in-development", CheckInDevelopment)

	// add state machines
	if err = state_machine.AddStateMachine(StateMachineA, "/config/state_machine_a.json"); err != nil {
		panic(err)
	}
	if err = state_machine.AddStateMachine(StateMachineB, "/config/state_machine_b.json"); err != nil {
		panic(err)
	}

	// check transitions of state machines
	stateMachines := []string{StateMachineA, StateMachineB}
	maxLen := 5
	ok := false

	for _, stateMachine := range stateMachines {
		fmt.Printf("\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err = state_machine.CheckTransition(stateMachine, i, j, 1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\ntransition from %d to %d ? %t", i, j, ok)
			}
		}
	}

    // Get all transitions
	transitions, err := state_machine.GetTransitions(StateMachineA, 1)
	for _, transition := range transitions {
		fmt.Printf("\ncan make transition to %s", transition.Name)
	}
}

func CheckInProgress(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecuting check in-progress handler with %+v", args)
	return true, nil
}

func CheckInDevelopment(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecuting check in-development handler with %+v", args)
	return true, nil
}
```

>### Result
```
State Machine: A

transition from 1 to 5 ? false
transition from 1 to 4 ? false
transition from 1 to 3 ? false
transition from 1 to 2 ? true
transition from 1 to 1 ? false
transition from 2 to 5 ? false
transition from 2 to 4 ? true
executing check in-progress handler with [[[1 text true]]]
transition from 2 to 3 ? true
transition from 2 to 2 ? false
transition from 2 to 1 ? false
transition from 3 to 5 ? false
transition from 3 to 4 ? false
transition from 3 to 3 ? false
transition from 3 to 2 ? false
transition from 3 to 1 ? false
transition from 4 to 5 ? false
transition from 4 to 4 ? false
transition from 4 to 3 ? false
transition from 4 to 2 ? false
transition from 4 to 1 ? false
transition from 5 to 5 ? false
transition from 5 to 4 ? false
transition from 5 to 3 ? false
transition from 5 to 2 ? false
transition from 5 to 1 ? false
State Machine: B

transition from 1 to 5 ? false
transition from 1 to 4 ? false
transition from 1 to 3 ? false
transition from 1 to 2 ? true
transition from 1 to 1 ? false
transition from 2 to 5 ? false
transition from 2 to 4 ? true
executing check in-development handler with [[[1 text true]]]
transition from 2 to 3 ? true
transition from 2 to 2 ? false
transition from 2 to 1 ? false
transition from 3 to 5 ? false
transition from 3 to 4 ? false
transition from 3 to 3 ? false
transition from 3 to 2 ? false
transition from 3 to 1 ? false
transition from 4 to 5 ? false
transition from 4 to 4 ? false
transition from 4 to 3 ? false
transition from 4 to 2 ? false
transition from 4 to 1 ? false
transition from 5 to 5 ? false
transition from 5 to 4 ? false
transition from 5 to 3 ? false
transition from 5 to 2 ? false
transition from 5 to 1 ? false
can make transition to In progress
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
