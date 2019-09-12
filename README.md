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
func main() {
	stateMachine, err := state_machine.New(
		state_machine.WithTransitionCheckHandler("check_3", Check3),
	)
	if err != nil {
		panic(err)
	}

	ok1, err := stateMachine.CheckTransition(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\ntransition from %d to %d ? %t", 1, 2, ok1)

	ok2, err := stateMachine.CheckTransition(2, 3, "1", 2, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\ntransition from %d to %d ? %t", 2, 3, ok2)

	ok3, err := stateMachine.CheckTransition(4, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\ntransition from %d to %d ? %t", 4, 1, ok3)

	ok4, err := stateMachine.CheckTransition(4, 5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\ntransition from %d to %d ? %t", 4, 5, ok4)
}

func Check3(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecuting check 3 handler with %+v", args)
	return true, nil
}
```

>### Result
```
transition from 1 to 2 ? true
executing check 3 handler with [[1 2 true]]
transition from 2 to 3 ? true
transition from 4 to 1 ? false
transition from 4 to 5 ? false
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
