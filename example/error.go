package main

import (
	"fmt"
	"go-log/app"
	"os"

	gowriter "github.com/joaosoft/go-writer/app"
)

func runTestAddition() {
	//
	// log to text
	fmt.Println(":: LOG TEXT")
	log := golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithFormatHandler(gowriter.JsonFormatHandler),
		golog.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME, "ip": golog.IP, "package": golog.PACKAGE, "function": golog.FUNCTION, "stack": golog.STACK, "trace": golog.TRACE},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	var err error
	log.Errorf("deu erro na linha %d", 201).ToError(&err)
	fmt.Printf("ERROR: %s", err.Error())
}
