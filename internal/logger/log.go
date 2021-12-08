package logger

import (
	"GitfyBot/internal"
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

func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

func setLogWriter() {
	var err error
	logWriter, err = os.OpenFile(internal.Config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	dropErr(err)
	log.SetOutput(io.MultiWriter(logWriter, os.Stdout))
}

func Info(args string) {
	setLogWriter()
	defer func(logWriter *os.File) {
		err := logWriter.Close()
		dropErr(err)
	}(logWriter)
	log.Info(args)
}

func Error(args string) {
	setLogWriter()
	defer func(logWriter *os.File) {
		err := logWriter.Close()
		dropErr(err)
	}(logWriter)
	log.Error(args)
}

func Fatal(args string) {
	setLogWriter()
	defer func(logWriter *os.File) {
		err := logWriter.Close()
		dropErr(err)
	}(logWriter)
	log.Fatal(args)
}
