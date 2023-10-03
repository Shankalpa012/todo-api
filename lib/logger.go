package lib

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLog *Logger
var zapLogger *zap.Logger

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

// GetLogger gets the global instance of the logger
func GetLogger() Logger {
	if globalLog != nil {
		return *globalLog
	}
	globalLog := newLogger()
	return *globalLog
}

// newLogger sets up logger the main logger
func newLogger() *Logger {

	env := "local"
	logLevel := "info"

	config := zap.NewDevelopmentConfig()

	if env == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}
	config.Level.SetLevel(level)
	zapLogger, _ = config.Build()

	globalLog := zapLogger.Sugar()

	return &Logger{
		SugaredLogger: globalLog,
	}

}
