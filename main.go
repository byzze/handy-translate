package main

import (
	"handy-translate/config"
	"handy-translate/log"
	"handy-translate/mywindow"
	"handy-translate/register"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file := log.Init()
	if file != nil {
		defer file.Close()
		file.Seek(0, 0)
		mw := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(mw)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	config.Init()
	go register.Hook()

	mywindow.Init()
}
