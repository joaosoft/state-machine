package logger

import (
	"fmt"
	"testing"
)

func TestNoneLogger(t *testing.T) {
	fmt.Println("test with NONE level")
	log := NewLogDefault("test", NoneLevel)
	log.Panic("panic")
	log.Info("info")
	log.Debug("debug")
}
