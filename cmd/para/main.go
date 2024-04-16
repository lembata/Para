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
	server, err := api.Init()
	if err != nil {
		logger.Errorf("failed to initialize server: %v", err)
		exitCode = 1
		return
	}

	err = server.Start()

	if err != nil {
		logger.Errorf("failed to start server: %v", err)
		exitCode = 1
		return
	}
	//defer server.Shutdown(context.WithCancel(context.Background(), os.Exit(exitCode))b
}

func recoverPanic() {
	if err := recover(); err != nil {
		exitCode = 1
		logger.Errorf("panic: %v\n%s", err, debug.Stack())
	}
}
