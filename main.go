package main

import (
	"io"
	"lyzee-translate/log"
	"lyzee-translate/mywindow"
	"lyzee-translate/register"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file := log.Init()
	if file != nil {
		defer file.Close()
		file.Seek(0, 0) // 每次运行清空日志
		mw := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(mw)
	}

	go register.Hook()

	mywindow.Init()
}
