package logger

import (
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/config"
	"fmt"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogWriter() zapcore.WriteSyncer {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.MkdirAll("logs", 0777)
		if err != nil {
			panic(err)
		}
	}

	// check and create file
	t := time.Now()
	var f *os.File
	logFile := fmt.Sprintf("logs/%v.log", t.Format("20060102"))
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ = os.Create(logFile)
	}
	f, _ = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	return zapcore.AddSync(f)
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// New is func initialize
func New(cfg *config.Log, e *gin.Engine) zap.Logger {
	cores := []zapcore.Core{}

	if cfg.EnableConsole {
		level := getZapLevel(cfg.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(cfg.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if cfg.EnableFile {
		level := getZapLevel(cfg.FileLevel)
		writer := getLogWriter()
		core := zapcore.NewCore(getEncoder(cfg.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	e.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	e.Use(ginzap.RecoveryWithZap(logger, true))

	e.GET("/ping", func(c *gin.Context) {
		c.String(200, "xxx")
	})

	e.GET("/huhu", func(c *gin.Context) {
		panic("s")
		// c.Error(err)
	})
	return *logger
}
