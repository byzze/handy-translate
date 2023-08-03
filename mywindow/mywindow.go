package mywindow

import (
	"handy-translate/config"
	"handy-translate/mywindow/fyne"
	"handy-translate/mywindow/lorca"
	"handy-translate/mywindow/webview"
)

// MyWindow interface
type MyWindow interface {
	Run()
}

type webviewWindow struct{}

type fyneWindow struct{}

type lorcaWindow struct{}

func (w *webviewWindow) Run() {
	webview.Run()
}

func (w *fyneWindow) Run() {
	fyne.Run()
}

func (w *lorcaWindow) Run() {
	lorca.Run()
}

// Init windows 初始化
func Init() {
	// newMyWindow("webview").Run()
	// newMyWindow("fyne").Run()
	// newMyWindow("lorca").Run()
	newMyWindow(config.Data.WindowWay).Run()
}

func newMyWindow(w string) MyWindow {
	switch w {
	case "fyne":
		return new(fyneWindow)
	case "webview":
		return new(webviewWindow)
	case "lorca":
		return new(lorcaWindow)
	}
	return nil
}
