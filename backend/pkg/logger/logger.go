package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(loggerLevel string) {
	// Логгирование
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, err := os.OpenFile("logs/golog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка создания файла golog", err)
	}
	writer := zapcore.AddSync(logFile)

	var defaultLogLevel zapcore.Level
	switch loggerLevel {
	case "debug":
		defaultLogLevel = zapcore.DebugLevel
	case "warn":
		defaultLogLevel = zapcore.WarnLevel
	case "error":
		defaultLogLevel = zapcore.ErrorLevel
	case "fatal":
		defaultLogLevel = zapcore.FatalLevel
	case "info":
	default:
		defaultLogLevel = zapcore.InfoLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}
