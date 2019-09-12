package main

import (
	"fmt"
	"state-machine"
)

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

	ok2, err := stateMachine.CheckTransition(2, 3, "1", 2, true)
	if err != nil {
		panic(err)
	}

	ok3, err := stateMachine.CheckTransition(4, 1)
	if err != nil {
		panic(err)
	}

	ok4, err := stateMachine.CheckTransition(4, 5)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ntransition from %d to %d ? %t", 1, 2, ok1)
	fmt.Printf("\ntransition from %d to %d ? %t", 2, 3, ok2)
	fmt.Printf("\ntransition from %d to %d ? %t", 4, 1, ok3)
	fmt.Printf("\ntransition from %d to %d ? %t", 4, 5, ok4)
}

func Check3(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecuting check 3 handler with %+v", args)
	return true, nil
}
