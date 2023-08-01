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

func init() {
	config.Init()
}

func main() {
	file := log.Init()
	if file != nil {
		defer file.Close()
		file.Seek(0, 0) // 每次运行清空日志
		mw := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(mw)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	go register.Hook()

	mywindow.Init()
}
