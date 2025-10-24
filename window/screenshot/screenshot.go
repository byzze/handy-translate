package screenshot

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log/slog"

	"github.com/kbinani/screenshot"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

var ScreenshotImg *image.RGBA

var WindowName = "Screenshot"

var Window *application.WebviewWindow

// NewWindow 截图功能也可以提取成一个单独程序，设计screenshot，robotgo库的使用
func NewWindow(app *application.App) {
	Window = app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:           WindowName,
		InitialPosition: application.WindowCentered,
		Hidden:          true,
		KeyBindings: map[string]func(window application.Window){
			"escape": func(window application.Window) {
				window.Hide()
			},
			"F12": func(window application.Window) {
				if w, ok := window.(*application.WebviewWindow); ok {
					w.OpenDevTools()
				}
			},
		},
		BackgroundType: application.BackgroundTypeTransparent,
		URL:            "http://wails.localhost/screenshot.html",
	})

	Window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowClosing win2")
		e.Cancel()
		Window.Hide()
	})

}

func ScreenshotFullScreen() string {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)

	if err != nil {
		// 错误处理，输出错误信息并返回
		slog.Error("Error capturing screenshot:", err)
		return ""
	}
	ScreenshotImg = img

	base64Image := encodeImageToBase64(img)
	if base64Image == "" {
		// 错误处理，未能生成Base64图像，返回
		slog.Error("Error encoding image to Base64")
		return ""
	}
	return base64Image

}

// CaptureSelectedScreen 截图功能
func CaptureSelectedScreen(startX, startY, endwidth, endheight int) image.Image {
	slog.Info("CaptureSelectedScreen",
		slog.Any("startX", startX),
		slog.Any("startY", startY),
		slog.Any("endwidth", endwidth),
		slog.Any("endheight", endheight))

	// x, y, width, height := startX, startY, endwidth, endheight
	// 裁剪图片
	/* 	croppedImg, err := screenshot.Capture(x, y, width, height)
	   	if err != nil {
	   		slog.Error("Capture", err)
	   		return nil
	   	} */

	// 裁剪图片
	rect := image.Rect(startX, startY, endwidth, endheight)
	if ScreenshotImg == nil {
		bounds := screenshot.GetDisplayBounds(0)
		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			// 错误处理，输出错误信息并返回
			fmt.Println("Error capturing screenshot:", err)
			return nil
		}
		ScreenshotImg = img
	}

	croppedImg := ScreenshotImg.SubImage(rect)

	return croppedImg
}

// 将图像编码为Base64字符串
func encodeImageToBase64(img image.Image) string {
	// 创建一个缓冲区用于保存Base64编码的数据
	var imgBytes []byte
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		panic(err)
	}

	imgBytes = buf.Bytes()

	// 使用base64编码图像数据
	base64Image := base64.StdEncoding.EncodeToString(imgBytes)

	return base64Image
}
