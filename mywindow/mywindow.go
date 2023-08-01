package mywindow

import (
	"lyzee-translate/mywindow/fyne"
	"lyzee-translate/mywindow/lorca"
	"lyzee-translate/mywindow/webview"
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
	newMyWindow("lorca").Run()
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
