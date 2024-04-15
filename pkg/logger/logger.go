package log

import (
	"fmt"
	"path/filepath"
	"os"
	"time"
	"github.com/rs/zerolog"
)

type Logger struct {
	logger         *zerolog.Logger// logrus.Logger
}

func NewLogger() *Logger {
	zerologger := zerolog.New(zerolog.
			ConsoleWriter{
				Out: os.Stderr,
				TimeFormat: time.RFC3339,
				FormatCaller: func(i interface{}) string {
					return "[" + filepath.Base(fmt.Sprintf("%s", i)) + "]"
				},
			}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			CallerWithSkipFrameCount(3).
			Logger();

	logger := &Logger{ 
		logger: &zerologger,
	}

	return logger
}

func (log *Logger) WithField(key string, value string) *Logger {
	new_logger := log.logger.
		With().
		Str(key, value).
		Logger()

	log.logger = &new_logger

	return log
}

func (log *Logger) Trace(msg string) {
	log.logger.Trace().Msg(msg)
}

func (log *Logger) Tracef(format string, msg ...interface{}) {
	log.logger.Trace().Msgf(format, msg...)
}

func (log *Logger) Debug(msg string) {
	log.logger.Debug().Msg(msg)
}

func (log *Logger) Debugf(format string, msg ...interface{}) {
	log.logger.Debug().Msgf(format, msg...)
}

func (log *Logger) Info(msg string) {
	log.logger.Info().Msg(msg)
}

func (log *Logger) Infof(format string, msg ...interface{}) {
	log.logger.Info().Msgf(format, msg...)
}

func (log *Logger) Warn(msg string) {
	log.logger.Warn().Msg(msg)
}

func (log *Logger) Warnf(format string, msg ...interface{}) {
	log.logger.Warn().Msgf(format, msg...)
}

func (log *Logger) Error(msg string) {
	log.logger.Error().Msg(msg)
}

func (log *Logger) Errorf(format string, msg ...interface{}) {
	log.logger.Error().Msgf(format, msg...)
}

func (log *Logger) Fatal(msg string) {
	log.logger.Fatal().Msg(msg)
}

func (log *Logger) Fatalf(format string, msg ...interface{}) {
	log.logger.Fatal().Msgf(format, msg...)
}
