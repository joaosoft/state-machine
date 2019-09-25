package main

import "fmt"

func CheckInProgress(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress handler with %+v", args)
	return false, nil
}

func ExecuteInProgress(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-progress handler with %+v", args)
	return true, nil
}

func EventOnSuccessInProgress(args ...interface{}) error {
	fmt.Printf("\nsuccess event in-progress handler with %+v", args)
	return nil
}

func EventOnErrorInProgress(args ...interface{}) error {
	fmt.Printf("\nerror event in-progress handler with %+v", args)
	return nil
}

func CheckInDevelopment(args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development handler with %+v", args)
	return true, nil
}

func ExecuteInDevelopment(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecute in-development handler with %+v", args)
	return true, nil
}

func EventOnSuccessInDevelopment(args ...interface{}) error {
	fmt.Printf("\nsuccess event in-development handler with %+v", args)
	return nil
}

func EventOnErrorInDevelopment(args ...interface{}) error {
	fmt.Printf("\nerror event in-development handler with %+v", args)
	return nil
}
