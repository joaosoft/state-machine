package main

import "fmt"

// ::  state machine A
// check
func CheckNewToInProgress(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress handler with %+v", args)
	return true, nil
}

func CheckInProgressToApproved(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress to approved handler with %+v", args)
	return true, nil
}

func CheckInProgressToDenied(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress to denied handler with %+v", args)
	return true, nil
}

// execute
func ExecuteNewToInProgress(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-progress handler with %+v", args)
	return true, nil
}

func ExecuteInProgressToApproved(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-progress to approved handler with %+v", args)
	return true, nil
}

func ExecuteInProgressToDenied(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-progress to denied handler with %+v", args)
	return true, nil
}

// event success
func EventOnSuccessNewToInProgress(args ...interface{}) error {
	fmt.Printf("\nsuccess event in-progress handler with %+v", args)
	return nil
}

func EventOnSuccessInProgressToApproved(args ...interface{}) error {
	fmt.Printf("\nsuccess event approved handler with %+v", args)
	return nil
}

func EventOnSuccessInProgressToDenied(args ...interface{}) error {
	fmt.Printf("\nsuccess event denied handler with %+v", args)
	return nil
}

// event error
func EventOnErrorNewToInProgress(args ...interface{}) error {
	fmt.Printf("\nerror event in-progress handler with %+v", args)
	return nil
}

func EventOnErrorInProgressToApproved(args ...interface{}) error {
	fmt.Printf("\nerror event approved handler with %+v", args)
	return nil
}

func EventOnErrorInProgressToDenied(args ...interface{}) error {
	fmt.Printf("\nerror event denied handler with %+v", args)
	return nil
}

// :: state machine B
// check
func CheckTodoToInDevelopment(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development handler with %+v", args)
	return true, nil
}

func CheckInDevelopmentToDone(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development to done handler with %+v", args)
	return true, nil
}

func CheckInDevelopmentToCanceled(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development to canceled handler with %+v", args)
	return true, nil
}

// execute
func ExecuteTodoToInDevelopment(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-development handler with %+v", args)
	return true, nil
}

func ExecuteInDevelopmentToCanceled(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute canceled handler with %+v", args)
	return true, nil
}

func ExecuteInDevelopmentToDone(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-progress to done handler with %+v", args)
	return true, nil
}

// event success
func EventOnSuccessInDevelopmentToDone(args ...interface{}) error {
	fmt.Printf("\nsuccess event done handler with %+v", args)
	return nil
}

func EventOnSuccessInDevelopmentToCanceled(args ...interface{}) error {
	fmt.Printf("\nsuccess event canceled handler with %+v", args)
	return nil
}

func EventOnSuccessTodoToInDevelopment(args ...interface{}) error {
	fmt.Printf("\nsuccess event in-development handler with %+v", args)
	return nil
}

// event error
func EventOnErrorInDevelopmentToDone(args ...interface{}) error {
	fmt.Printf("\nerror event done handler with %+v", args)
	return nil
}

func EventOnErrorInDevelopmentToCanceled(args ...interface{}) error {
	fmt.Printf("\nerror event canceled handler with %+v", args)
	return nil
}

func EventOnErrorTodoToInDevelopment(args ...interface{}) error {
	fmt.Printf("\nerror event in-development handler with %+v", args)
	return nil
}
