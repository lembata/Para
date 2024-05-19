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
	defer recoverPanic()
	defer func() {
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	logger.Info("Welcome to Para!")
	configDir := getDefaultConfigDir()

	logger.Infof("Using config directory: %s", configDir)

	db, err := initDatabase(configDir)

	if err != nil {
		exitCode = 1
		logger.Errorf("failed to open database: %v", err)
		return
	}

	defer func(database *database.Database) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	server, err := api.Init()

	if err != nil {
		logger.Errorf("failed to init server: %v", err)
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

func initDatabase(configDir string) (*database.Database, error) {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0755); err != nil {
			return nil, err
		}
	}

	dbPath := path.Join(configDir, "para.sqlite")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if _, err := os.Create(dbPath); err != nil {
			return nil, err
		}
	}

	dbInst := database.Init()
	err := dbInst.Open(dbPath)

	return dbInst, err
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
