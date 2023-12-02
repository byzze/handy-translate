package screenshot

import (
	"image"
	"log/slog"

	"github.com/kbinani/screenshot"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// NewWindow 截图功能也可以提取成一个单独程序，设计screenshot，robotgo库的使用
func NewWindow(app *application.App) *application.WebviewWindow {
	win3 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:    "Screenshot",
		Centered: true,
		Hidden:   true,
		KeyBindings: map[string]func(window *application.WebviewWindow){
			"escape": func(window *application.WebviewWindow) {
				window.Hide()
			},
		},
		BackgroundType: application.BackgroundTypeTransparent,
		URL:            "screenshot.html",
	})
	return win3
}

// CaptureSelectedScreen 截图功能
func CaptureSelectedScreen(startX, startY, endwidth, endheight int) *image.RGBA {
	slog.Info("CaptureSelectedScreen",
		slog.Any("startX", startX),
		slog.Any("startY", startY),
		slog.Any("endwidth", endwidth),
		slog.Any("endheight", endheight))

	x, y, width, height := startX, startY, endwidth, endheight
	// 裁剪图片
	croppedImg, err := screenshot.Capture(x, y, width, height)
	if err != nil {
		slog.Error("Capture", err)
		return nil
	}
	return croppedImg
}
