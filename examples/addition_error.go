package main

import (
	"fmt"
	logger "../../logger"
	"os"

	writer "github.com/joaosoft/writers"
)

func ExampleAdditionError() {
	//
	// log to text
	fmt.Println(":: LOG TEXT")
	log := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(writer.TextFormatHandler),
		logger.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": logger.LEVEL, "timestamp": logger.TIMESTAMP, "date": logger.DATE, "time": logger.TIME, "ip": logger.IP, "package": logger.PACKAGE, "function": logger.FUNCTION, "stack": logger.STACK, "trace": logger.TRACE},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"},
			map[string]interface{}{"ip": logger.IP, "function": logger.FUNCTION, "file": logger.FILE})

	var err error
	log.Errorf("deu erro na linha %d", 201).ToError(&err)
	fmt.Printf("ERROR: %s", err.Error())
}
