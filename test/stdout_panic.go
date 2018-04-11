package main

import (
	"fmt"
	"go-log/service"
	"time"

	"github.com/joaosoft/go-writer/service"
)

func runTestStdoutPanic() {
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
	log := golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithSpecialWriter(stdoutWriter)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
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
