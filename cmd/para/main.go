package main

import (
	"os"
	"runtime/debug"
	"github.com/lembata/para/pkg/logger"
	"github.com/lembata/para/internal/api"
)

var exitCode = 0;
var logger = log.NewLogger() 

func main() {
	defer func() {
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	logger.Info("Welcome to Para!")
	api.Init()
}

func recoverPanic() {
	if err := recover(); err != nil {
		exitCode = 1
		logger.Errorf("panic: %v\n%s", err, debug.Stack())
	}
}
