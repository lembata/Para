package api

import (
	"github.com/lembata/para/pkg/logger"
)

var logger = log.NewLogger()

func Init() {
	logger.Debug("Initializing API...")
}
