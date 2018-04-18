package main

import (
	"fmt"
	"go-log/service"
	"os"

	"github.com/joaosoft/go-writer/service"
)

func runTestAddition() {
	//
	// log to text
	fmt.Println(":: LOG TEXT")
	log := golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithFormatHandler(gowriter.TextFormatHandler),
		golog.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	var err error
	log.Errorf("deu erro na linha %d", 201).ToError(&err)
	fmt.Printf("ERROR: %s", err.Error())
}
