package main

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

var logger *zap.SugaredLogger

func InitLogger() error {
	logType, ok := os.LookupEnv("LOG_TYPE")
	if !ok {
		logType = "JSON"
	}

	if strings.ToLower(logType) == "simple" {
		l, err := zap.NewDevelopment()
		logger = l.Sugar()
		return err
	}

	l, err := zap.NewProduction()
	logger = l.Sugar()
	return err
}

func log2() *zap.SugaredLogger {
	return logger
}
