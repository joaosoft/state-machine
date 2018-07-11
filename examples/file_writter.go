package main

import (
	"fmt"
	logger "../../logger"
	"time"

	writer "github.com/joaosoft/writers"
)

func ExampleFileWritter() {
	//
	// file fileWriter
	quit := make(chan bool)
	fileWriter := writer.NewFileWriter(
		writer.WithFileDirectory("./testing"),
		writer.WithFileName("dummy_"),
		writer.WithFileMaxMegaByteSize(1),
		writer.WithFileFlushTime(time.Second*5),
		writer.WithFileQuitChannel(quit),
	)

	//
	// log to json
	fmt.Println(":: LOG JSON")
	log := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithSpecialWriter(fileWriter)).
		With(
			map[string]interface{}{"level": logger.LEVEL, "timestamp": logger.TIMESTAMP, "date": logger.DATE, "time": logger.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"},
			map[string]interface{}{"ip": logger.IP, "function": logger.FUNCTION, "file": logger.FILE})

	// logging...
	start := time.Now()
	sum := 0
	for i := 0; i < 100000; i++ {
		log.Infof("MESSAGE %d", i+1)
		sum += 1
	}
	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())

	<-time.After(time.Second * 10)
	quit <- true
}
