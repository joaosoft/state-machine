# go-log
[![Build Status](https://travis-ci.org/joaosoft/go-log.svg?branch=master)](https://travis-ci.org/joaosoft/go-log) | [![codecov](https://codecov.io/gh/joaosoft/go-log/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/go-log) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/go-log)](https://goreportcard.com/report/github.com/joaosoft/go-log) | [![GoDoc](https://godoc.org/github.com/joaosoft/go-log?status.svg)](https://godoc.org/github.com/joaosoft/go-log/app)

A simplified logger that allows you to add complexity depending of your requirements.
The easy way to use the logger:
``` Go
import log github.com/joaosoft/go-log/app




log.Info("hello")
```
you also can config it, as i prefer, please see below
After a read of the project https://gitlab.com/vredens/go-logger extracted some concepts like allowing to add tags and fields to logger infrastructure. 

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* formatted messages
* prefixes (special prefixes: DATE, TIME, TIMESTAMP, LEVEL, IP, PACKAGE, FUNCTION, FILE, TRACE, STACK)
* tags
* fields
* writers at [[go-writer]](https://github.com/joaosoft/go-writer/tree/master/bin/example)
  * to file (with queue processing)[1] 
  * to stdout (with queue processing)[1] [[here]](https://github.com/joaosoft/go-writer/tree/master/example)
* addition commands (ToError(&err))
  
  [1] this writer allows you to continue the processing and dispatch the logging

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/go-log/service
```

## Interface 
```go
type Log interface {
	SetLevel(level Level)

	With(prefixes, tags, fields map[string]interface{}) ILog
	WithPrefixes(prefixes map[string]interface{}) ILog
	WithTags(tags map[string]interface{}) ILog
	WithFields(fields map[string]interface{}) ILog

	WithPrefix(key string, value interface{}) ILog
	WithTag(key string, value interface{}) ILog
	WithField(key string, value interface{}) ILog

	Debug(message interface{}) IAddition
	Info(message interface{}) IAddition
	Warn(message interface{}) IAddition
	Error(message interface{}) IAddition

	Debugf(format string, arguments ...interface{}) IAddition
	Infof(format string, arguments ...interface{}) IAddition
	Warnf(format string, arguments ...interface{}) IAddition
	Errorf(format string, arguments ...interface{}) IAddition
	
	Reconfigure(options ...logOption)
}

type IAddition interface {
	ToError(err *error) IAddition
	ToErrorData(err *goerror.ErrorData) IAddition
}

type ISpecialWriter interface {
	SWrite(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) (n int, err error)
}

```

## Usage 
This examples are available in the project at [go-log/example](https://github.com/joaosoft/go-log/tree/master/example)

```go
//
// log to text
fmt.Println(":: LOG TEXT")
log := golog.NewLog(
    golog.WithLevel(golog.InfoLevel), 
    golog.WithFormatHandler(golog.TextFormatHandler), 
    golog.WithWriter(os.Stdout)).
        With(
            map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
            map[string]interface{}{"service": "log"}, 
            map[string]interface{}{"name": "joão"})

// logging...
log.Error("isto é uma mensagem de error")
log.Info("isto é uma mensagem de info")
log.Debug("isto é uma mensagem de debug")

fmt.Println("--------------")
<-time.After(time.Second)

//
// log to json
fmt.Println(":: LOG JSON")
log = golog.NewLog(
    golog.WithLevel(golog.InfoLevel),
    golog.WithFormatHandler(golog.JsonFormatHandler),
    golog.WithWriter(os.Stdout)).
        With(
            map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
            map[string]interface{}{"service": "log"},
            map[string]interface{}{"name": "joão"})

// logging...
log.Errorf("isto é uma mensagem de error %s", "hello")
log.Infof("isto é uma  mensagem de info %s ", "hi")
log.Debugf("isto é uma mensagem de debug %s", "ehh")

// error...
var err error
log.Errorf("deu erro na linha %d", 201).ToError(&err)
fmt.Printf("ERROR: %s", err.Error())
```

###### Output 

```javascript
:: LOG TEXT
{prefixes:map[level:error time:2018-03-20 02:47:21] tags:map[service:log] message:isto é uma mensagem de error fields:map[name:joão]}
{prefixes:map[level:info time:2018-03-20 02:47:21] tags:map[service:log] message:isto é uma mensagem de info fields:map[name:joão]}
--------------
:: LOG JSON
{"prefixes":{"level":"error","time":"2018-03-20 02:47:22"},"tags":{"service":"log"},"message":"isto é uma mensagem de error hello","fields":{"name":"joão"}}
{"prefixes":{"level":"info","time":"2018-03-20 02:47:22"},"tags":{"service":"log"},"message":"isto é uma  mensagem de info hi ","fields":{"name":"joão"}}
```

## Known issues
* all the maps do not guarantee order of the items! 


## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
