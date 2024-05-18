package main

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime/debug"

	"github.com/lembata/para/internal/api"
	"github.com/lembata/para/pkg/database"
	"github.com/lembata/para/pkg/logger"
)

var exitCode = 0
var logger = log.NewLogger()

func main() {
	defer func() {
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	logger.Info("Welcome to Para!")
	configDir := getDefaultConfigDir()

	logger.Infof("Using config directory: %s", configDir)

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0755)
	}

	dbPath := path.Join(configDir, "para.sqlite")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Create(dbPath)
	}

	db := database.Init()
	defer db.Close()

	server, err := api.Init()

	if err != nil {
		logger.Errorf("failed to initialize server: %v", err)
		exitCode = 1
		return
	}

	database := database.Init()
	database.Open(dbPath)
	err = database.Open(dbPath)

	if err != nil {
		exitCode = 1
		logger.Errorf("failed to open database: %v", err)
		return
	}

	defer database.Close()

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

func getDefaultConfigDir() string {
	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	return filepath.Join(currentUser.HomeDir, ".config/para")
}
