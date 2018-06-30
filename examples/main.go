package main

import (
	"fmt"
)

func main() {
	// addition
	fmt.Println("\naddition...")
	ExampleAdditionError()

	// default
	fmt.Println("\ndefault...")
	ExampleDefaultLogger()

	// write log to file with queue
	fmt.Println("\nwrite to file...")
	ExampleFileWritter()

	// write log to stdout with queue
	fmt.Println("\nwrite to stdout...")
	ExampleStdoutWritter()

	// write log to stdout with queue on panic
	fmt.Println("\nwrite to stdout on panic...")
	ExampleStdoutWritterWithPanic()

	// write log to file with queue on panic
	fmt.Println("\nwrite to file on panic...")
	ExampleFileWritterWithPanic()
}
