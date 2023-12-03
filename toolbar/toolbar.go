package toolbar

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

var WindowName = "ToolBar"

var Window *application.WebviewWindow

// NewWindow 截图功能也可以提取成一个单独程序，设计screenshot，robotgo库的使用
func NewWindow(app *application.App) {
	Window = app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:         WindowName,
		Width:         300,
		Height:        54,
		AlwaysOnTop:   false,
		Hidden:        true,
		DisableResize: false,
		Frameless:     true,
		Centered:      true,
		ShouldClose: func(window *application.WebviewWindow) bool {
			app.Quit()
			return true
		},

		URL: "index.html",
	})
	Window.SetMaxSize(300, 600)
	Window.On(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		Window.Hide()
	})
}
