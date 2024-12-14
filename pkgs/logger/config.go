package logger

import (
	"go.uber.org/zap/zapcore"
)

// levelMappings is a map of string to zapcore.Level
// It is used to map the string value of the logFn level to the zapcore.Level
// type.
var levelMappings = map[string]zapcore.Level{ //nolint:gochecknoglobals
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

// Config is the configuration for the logger.
// It contains the logFn level.
// The logFn level can be one of the following:
// - debug
// - info
// - warn
// - error.
type Config struct {
	Level  string
	Writer zapcore.WriteSyncer
}

// GetLevel returns the zapcore.Level for the logFn level in the config.
// If the logFn level is not one of the supported values, it returns the
// default logFn level, which is info.
func (c Config) GetLevel() zapcore.Level {
	if level, ok := levelMappings[c.Level]; ok {
		return level
	}

	return zapcore.InfoLevel
}
