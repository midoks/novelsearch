package logs

import (
	"fmt"

	"github.com/midoks/novelsearch/internal/conf"
	go_logger "github.com/phachon/go-logger"
)

var (
	logFileName = "novelsearch.log"
	logger      *go_logger.Logger
)

func Init() *go_logger.Logger {

	logger = go_logger.NewLogger()

	jsonFormat := false
	if conf.Log.Format == "json" {
		jsonFormat = true
	}

	logPath := fmt.Sprintf("%s", conf.Log.RootPath)
	fileConfig := &go_logger.FileConfig{
		Filename: fmt.Sprintf("%s/%s", logPath, logFileName),
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): fmt.Sprintf("%s/%s", logPath, "error.log"),
			logger.LoggerLevel("info"):  fmt.Sprintf("%s/%s", logPath, "info.log"),
			logger.LoggerLevel("debug"): fmt.Sprintf("%s/%s", logPath, "debug.log"),
		},
		MaxSize:    1024 * 1024,
		MaxLine:    100000,
		DateSlice:  "d",
		JsonFormat: jsonFormat,
		Format:     "",
	}
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)

	return logger
}

func GetLogger() *go_logger.Logger {
	return logger
}

func Debug(args string) {
	logger.Debug(args)
}

func Info(args string) {
	logger.Info(args)
}

func Warn(args string) {
	logger.Warning(args)
}

func Error(args string) {
	logger.Error(args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
