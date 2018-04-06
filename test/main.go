package main

import (
	"fmt"
)

func main() {
	// write log to file with queue
	fmt.Println("write to file...")
	runTestFile()

	// write log to stdout with queue
	fmt.Println("write to stdout...")
	runTestStdout()

	// write log to stdout with queue on panic
	fmt.Println("write to stdout on panic...")
	runTestStdoutPanic()

	// write log to file with queue on panic
	fmt.Println("write to file on panic...")
	runTestFilePanic()
}
