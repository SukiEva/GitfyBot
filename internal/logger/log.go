package logger

import (
	"GitfyBot/internal/utils"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	log       = logrus.New()
	logWriter *os.File
)

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
}

func setLogWriter() {
	var err error
	logWriter, err = os.OpenFile(utils.Config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utils.DropErr(err)
	log.SetOutput(io.MultiWriter(logWriter, os.Stdout))
}

func Info(args string) {
	setLogWriter()
	defer func(logWriter *os.File) {
		err := logWriter.Close()
		utils.DropErr(err)
	}(logWriter)
	log.Info(args)
}

func Error(args string) {
	setLogWriter()
	defer func(logWriter *os.File) {
		err := logWriter.Close()
		utils.DropErr(err)
	}(logWriter)
	log.Error(args)
}
