package main

import (
	"fmt"
	"os"
	"time"

	"../../logger"

	writer "github.com/joaosoft/writers"
)

type Example struct{}

func main() {
	e := Example{}

	// default
	fmt.Println("\ndefault...")
	e.ExampleDefaultLogger()

	// addition
	fmt.Println("\naddition...")
	e.ExampleAdditionError()

	// write log to file with queue
	fmt.Println("\nwrite to file...")
	e.ExampleFileWritter()

	// write log to stdout with queue
	fmt.Println("\nwrite to stdout...")
	e.ExampleStdoutWritter()

	// write log to stdout with queue on panic
	fmt.Println("\nwrite to stdout on panic...")
	e.ExampleStdoutWritterWithPanic()

	// write log to file with queue on panic
	fmt.Println("\nwrite to file on panic...")
	e.ExampleFileWritterWithPanic()
}

func (e Example) ExampleAdditionError() {
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

	err := log.Errorf("deu erro na linha %d", 201).ToError()
	fmt.Printf("ERROR: %s", err.Error())
}

func (e Example) ExampleDefaultLogger() {
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

func (e Example) ExampleFileWritter() {
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

func (e Example) ExampleStdoutWritter() {
	//
	// stdout fileWriter
	quit := make(chan bool)
	stdoutWriter := writer.NewStdoutWriter(
		writer.WithStdoutFormatHandler(writer.JsonFormatHandler),
		writer.WithStdoutFlushTime(time.Second*5),
		writer.WithStdoutQuitChannel(quit),
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

func (e Example) ExampleStdoutWritterWithPanic() {
	//
	// stdout fileWriter
	quit := make(chan bool)
	stdoutWriter := writer.NewStdoutWriter(
		writer.WithStdoutFormatHandler(writer.JsonFormatHandler),
		writer.WithStdoutFlushTime(time.Second*5),
		writer.WithStdoutQuitChannel(quit),
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
			map[string]interface{}{"name": "joão"},
			map[string]interface{}{"ip": logger.IP, "function": logger.FUNCTION, "file": logger.FILE})

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

func (e Example) ExampleFileWritterWithPanic() {
	//
	// stdout fileWriter
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

		if i == 50000 {
			panic("FUCKED!")
		}
	}
	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())

	<-time.After(time.Second * 10)
	quit <- true
}
