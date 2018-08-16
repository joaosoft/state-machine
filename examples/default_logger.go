package main

import (
	"fmt"
	"os"
	"time"

	logger "../../logger"

	writer "github.com/joaosoft/writers"
)

func ExampleDefaultLogger() {
	//
	// log to text
	fmt.Println(":: LOG TEXT")
	log := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(writer.TextFormatHandler),
		logger.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": logger.LEVEL, "timestamp": logger.TIMESTAMP, "date": logger.DATE, "time": logger.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"},
			map[string]interface{}{"ip": logger.IP, "function": logger.FUNCTION, "file": logger.FILE})

	// logging...
	log.Error("isto é uma mensagem de error")
	log.Info("isto é uma mensagem de info")
	log.Debug("isto é uma mensagem de debug")
	log.Error("")

	fmt.Println("--------------")
	<-time.After(time.Second)

	//
	// log to json
	fmt.Println(":: LOG JSON")
	log = logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(writer.JsonFormatHandler),
		logger.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": logger.LEVEL, "timestamp": logger.TIMESTAMP, "date": logger.DATE, "time": logger.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"},
			map[string]interface{}{"ip": logger.IP, "function": logger.FUNCTION, "file": logger.FILE})

	// logging...
	log.Errorf("isto é uma mensagem de error %s", "hello")
	log.Infof("isto é uma  mensagem de info %s ", "hi")
	log.Debugf("isto é uma mensagem de debug %s", "ehh")
}
