package mywindow

import (
	"lyzee-translate/mywindow/fyne"
	"lyzee-translate/mywindow/lorca"
	"lyzee-translate/mywindow/walk"
	"lyzee-translate/mywindow/webview"
)

type MyWindow interface {
	Run()
}

type webviewWindow struct{}

type fyneWindow struct{}

type walkWindow struct{}

func (w *webviewWindow) Run() {
	webview.Run()
}

func (w *fyneWindow) Run() {
	fyne.Run()
}
func (w *walkWindow) Run() {
	walk.Run()
}

func Init() {
	lorca.Run()
	// newMyWindow("webview").Run()
	// newMyWindow("fyne").Run()
	// newMyWindow("walk").Run()
}

func newMyWindow(w string) MyWindow {
	switch w {
	case "fyne":
		return new(fyneWindow)
	case "webview":
		return new(webviewWindow)
	case "walk":
		return new(walkWindow)
	}
	return nil
}
