package main

import (
	"fmt"
	logger "logger/models"
	"time"

	gowriter "github.com/joaosoft/go-writer/app"
)

func ExampleStdoutWritterWithPanic() {
	//
	// stdout fileWriter
	quit := make(chan bool)
	stdoutWriter := gowriter.NewStdoutWriter(
		gowriter.WithStdoutFormatHandler(gowriter.JsonFormatHandler),
		gowriter.WithStdoutFlushTime(time.Second*5),
		gowriter.WithStdoutQuitChannel(quit),
	)

	//
	// log to json
	fmt.Println(":: LOG JSON")
	log := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithSpecialWriter(stdoutWriter)).
		With(
			map[string]interface{}{"level": logger.LEVEL, "timestamp": logger.TIMESTAMP, "date": logger.DATE, "time": logger.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	// logging...
	start := time.Now()
	sum := 0
	for i := 0; i < 100000; i++ {
		log.Infof("MESSAGE %d", i+1)
		sum += 1

		if i == 50000 {
			panic("FUCKED!")
		}
	}
	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())

	<-time.After(time.Second * 10)
	quit <- true
}
