package translate

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

var WindowName = "Translate"

var Window *application.WebviewWindow

// NewWindow 截图功能也可以提取成一个单独程序，设计screenshot，robotgo库的使用
func NewWindow(app *application.App) {
	Window = app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     WindowName,
		Width:     500,
		Height:    500,
		Frameless: true,
		Hidden:    true,
		URL:       "translate.html",
	})

	Window.On(events.Common.WindowClosing, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowClosing win2")
		Window.Hide()
	})
}
