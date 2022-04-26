package log

import (
	"os"

	"github.com/mhd7966/arvan/wallet/configs"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LogInit() {

	Log = logrus.New()
	config := configs.Cfg.Log

	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-Jan-02 15:04:05",
	})

	if config.OutputType == "stdout" {
		Log.SetOutput(os.Stdout)

	} else if config.OutputType == "file" {
		file, err := os.OpenFile(config.OutputAdd, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			Log.Fatal(err)
		}
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)

	}

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	Log.SetLevel(logLevel)

}
