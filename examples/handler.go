package main

import (
	"fmt"
	state_machine "state-machine"
)

// ::  state machine A
// load
func loadDummy(ctx *state_machine.Context) error {
	fmt.Printf("load dummy handler with %+v\n", ctx.Args)
	return nil
}

// check
func checkNewToInProgress(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("check in-progress handler with %+v\n", ctx.Args)
	return true, nil
}

func checkInProgressToApproved(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("check in-progress to approved handler with %+v\n", ctx.Args)
	return true, nil
}

func checkInProgressToDenied(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("check in-progress to denied handler with %+v\n", ctx.Args)
	return true, nil
}

// execute
func executeNewToInProgress(ctx *state_machine.Context) error {
	fmt.Printf("execute in-progress handler with %+v\n", ctx.Args)
	return nil
}

func executeNewToInProgressByRole(ctx *state_machine.Context) error {
	fmt.Printf("by role: execute in-progress handler with %+v\n", ctx.Args)
	return nil
}

func executeInProgressToApproved(ctx *state_machine.Context) error {
	fmt.Printf("execute in-progress to approved handler with %+v\n", ctx.Args)
	return nil
}

func executeInProgressToDenied(ctx *state_machine.Context) error {
	fmt.Printf("execute in-progress to denied handler with %+v\n", ctx.Args)
	return nil
}

// event success
func eventOnSuccessNewToInProgress(ctx *state_machine.Context) error {
	fmt.Printf("success event in-progress handler with %+v\n", ctx.Args)
	return nil
}

func eventOnSuccessNewToInProgressByRole(ctx *state_machine.Context) error {
	fmt.Printf("by role: success event in-progress handler with %+v\n", ctx.Args)
	return nil
}

func eventOnSuccessInProgressToApproved(ctx *state_machine.Context) error {
	fmt.Printf("success event approved handler with %+v\n", ctx.Args)
	return nil
}

func eventOnSuccessInProgressToDenied(ctx *state_machine.Context) error {
	fmt.Printf("success event denied handler with %+v\n", ctx.Args)
	return nil
}

// event error
func eventOnErrorNewToInProgress(ctx *state_machine.Context, err error) error {
	fmt.Printf("error %s, event in-progress handler with %+v\n", err, ctx.Args)
	return nil
}

func eventOnErrorInProgressToApproved(ctx *state_machine.Context, err error) error {
	fmt.Printf("error %s, event approved handler with %+v\n", err, ctx.Args)
	return nil
}

func eventOnErrorInProgressToDenied(ctx *state_machine.Context, err error) error {
	fmt.Printf("error %s, event denied handler with %+v\n", err, ctx.Args)
	return nil
}

// transition handler
func StateMachineATransitionHandler(ctx *state_machine.Context) error {
	fmt.Printf("state machine: %s, transition handler with %+v\n", ctx.StateMachine, ctx.Args)
	return nil
}

// :: state machine B
// manual - init
func beforeExecuteLoadFromState(ctx *state_machine.Context) error {
	fmt.Printf("load 'from' state handler with %+v\n", ctx.Args)
	ctx.From = 1
	return nil
}

// check
func checkTodoToInDevelopment(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("check in-development handler with %+v\n", ctx.Args)
	return true, nil
}

func checkInDevelopmentToDone(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("check in-development to done handler with %+v\n", ctx.Args)
	return true, nil
}

func checkInDevelopmentToCanceled(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("check in-development to canceled handler with %+v\n", ctx.Args)
	return true, nil
}

// execute
func executeTodoToInDevelopment(ctx *state_machine.Context) error {
	fmt.Printf("execute in-development handler with %+v\n", ctx.Args)
	return nil
}

func executeInDevelopmentToCanceled(ctx *state_machine.Context) error {
	fmt.Printf("execute canceled handler with %+v\n", ctx.Args)
	return nil
}

func executeInDevelopmentToDone(ctx *state_machine.Context) error {
	fmt.Printf("execute in-progress to done handler with %+v\n", ctx.Args)
	return nil
}

// event success
func eventOnSuccessInDevelopmentToDone(ctx *state_machine.Context) error {
	fmt.Printf("success event done handler with %+v\n", ctx.Args)
	return nil
}

func eventOnSuccessInDevelopmentToCanceled(ctx *state_machine.Context) error {
	fmt.Printf("success event canceled handler with %+v\n", ctx.Args)
	return nil
}

func eventOnSuccessTodoToInDevelopment(ctx *state_machine.Context) error {
	fmt.Printf("success event in-development handler with %+v\n", ctx.Args)
	return nil
}

// event error
func eventOnErrorInDevelopmentToDone(ctx *state_machine.Context, err error) error {
	fmt.Printf("error %s, event done handler with %+v\n", err, ctx.Args)
	return nil
}

func eventOnErrorInDevelopmentToCanceled(ctx *state_machine.Context, err error) error {
	fmt.Printf("error %s, event canceled handler with %+v\n", err, ctx.Args)
	return nil
}

func eventOnErrorTodoToInDevelopment(ctx *state_machine.Context, err error) error {
	fmt.Printf("error %s, event in-development handler with %+v\n", err, ctx.Args)
	return nil
}

// transition handler
func StateMachineBTransitionHandler(ctx *state_machine.Context) error {
	fmt.Printf("state machine: %s, transition handler with %+v\n", ctx.StateMachine, ctx.Args)
	return nil
}
