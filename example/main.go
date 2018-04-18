package main

import (
	"fmt"
)

func main() {
	// addition
	fmt.Println("\naddition...")
	runTestAddition()

	// default
	fmt.Println("\ndefault...")
	runTestDefault()

	// write log to file with queue
	fmt.Println("\nwrite to file...")
	runTestFile()

	// write log to stdout with queue
	fmt.Println("\nwrite to stdout...")
	runTestStdout()

	// write log to stdout with queue on panic
	fmt.Println("\nwrite to stdout on panic...")
	runTestStdoutPanic()

	// write log to file with queue on panic
	fmt.Println("\nwrite to file on panic...")
	runTestFilePanic()
}
