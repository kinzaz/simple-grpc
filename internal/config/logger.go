package config

import (
	"errors"
	"os"
)

const (
	loggerLevelEnvName = "LEVEL_LOGGER"
)

type loggerConfig struct {
	level string
}

type LoggerConfig interface {
	Level() string
}

func NewLoggerConfig() (LoggerConfig, error) {
	level := os.Getenv(loggerLevelEnvName)
	if len(level) == 0 {
		return nil, errors.New("level not found")
	}

	return &loggerConfig{
		level: level,
	}, nil
}

func (cfg *loggerConfig) Level() string {
	return cfg.level
}
