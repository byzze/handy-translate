package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"handy-translate/config"
	"handy-translate/os_api/windows"
	"handy-translate/screenshot"
	"handy-translate/toolbar"
	"handy-translate/translate"
	"handy-translate/utils"
	"image/png"
	"log/slog"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// 和js绑定的go方法集合

// App is a service
type App struct {
}

// MyFetch URl
func (a *App) MyFetch(URL string, content map[string]interface{}) interface{} {
	return utils.MyFetch(URL, content)
}

// Transalte 翻译逻辑
func (a *App) Transalte(queryText, fromLang, toLang string) {
	app.Logger.Info("Transalte",
		slog.Any("queryText", queryText),
		slog.Any("toLang", toLang),
		slog.Any("fromLang", fromLang))

	processTranslate(queryText)
}

// GetTransalteMap 获取所有翻译配置
func (a *App) GetTransalteMap() string {
	var translateList = config.Data.Translate
	bTranslate, err := json.Marshal(translateList)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}
	return string(bTranslate)
}

// SetTransalteWay 设置当前翻译服务
func (a *App) SetTransalteWay(translateWay string) {
	config.Data.TranslateWay = translateWay
	translate.SetQueryText("")
	config.Save()
	slog.Info("SetTransalteList", slog.Any("config.Data.Translate", config.Data.Translate))
}

// GetTransalteWay 获取当前翻译的服务
func (a *App) GetTransalteWay() string {
	return config.Data.TranslateWay
}

// Show 通过名字控制窗口事件
func (a *App) Show(windowName string) {
	var win *application.WebviewWindow
	switch windowName {
	case screenshot.WindowName:
		win = screenshot.Window
	case translate.WindowName:
		win = translate.Window
	}
	win.Center()
	win.Show()
}

// Hide 通过名字控制窗口事件
func (a *App) Hide(windowName string) {
	var win *application.WebviewWindow
	switch windowName {
	case screenshot.WindowName:
		win = screenshot.Window
	case translate.WindowName:
		win = translate.Window
	}
	win.Hide()
}

// ToolBarShow 显示工具弹窗，控制大小，布局
func (a *App) ToolBarShow(height float64) {
	// 40 + 55 窗口空白区域+翻译的图标区域 + 44 + 55
	height = height + 40 + 50
	app.Logger.Info("ToolBarShow", slog.Float64("height", height))

	h := int(height)
	if h > 600 {
		h = 600
	}

	if h == 0 {
		h = 55
	}

	pos := windows.GetCursorPos()
	w := toolbar.Window
	w.SetSize(300, h)
	// x, y := robotgo.Location() // 在联想小新pro13 2k屏幕时数据不对
	x, y := int(pos.X), int(pos.Y) // 处理获取坐标不正确，采用windows原生api
	sc, _ := w.GetScreen()
	robotgo.GetMousePos()
	slog.Info("GetCursorPos", slog.Any("pos.X", pos.X), slog.Any("pos.Y", pos.Y))
	slog.Info("GetScreen", slog.Any("sc.Size.Width", sc.Size.Width), slog.Any("sc.Size.Height", sc.Size.Height))

	c := int(float64(sc.Size.Height) * 0.1)
	slog.Info("sc.Size.Height", slog.Any("c", c))
	if y+h+c >= sc.Size.Height {
		gap := y + h + c - sc.Size.Height
		slog.Info(">>>>", slog.Any("gap", gap))
		w.SetAbsolutePosition(x+10, y-gap)
	} else {
		slog.Info("<<<<")
		w.SetAbsolutePosition(x+10, y+10)
	}
	windows.FindWindow(toolbar.WindowName).ShowForWindows() // 使用原生showwindow，wails3版本有些问题，无法正常显示
	// windows.ShowForWindows(toolbar.WindowName) // 使用原生showwindow，wails3版本有些问题，无法正常显示
}

// CaptureSelectedScreen 截取选中的区域
func (a *App) CaptureSelectedScreen(startX, startY, width, height float64) {

	croppedImg := screenshot.CaptureSelectedScreen(int(startX), int(startY), int(width), int(height))
	if croppedImg == nil {
		return
	}

	var buf bytes.Buffer
	err := png.Encode(&buf, croppedImg)
	if err != nil {
		slog.Error("png.Encode", err)
		return
	}

	filename := "screenshot.png" // 保存的文件名
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	err = saveBase64Image(base64String, filename)
	if err != nil {
		logrus.Fatal("保存图片出错: ", err)
	}

	// OCR解析文本
	ocrres := ExecOCR(".\\RapidOCR-json.exe", filename)
	// 翻译文本
	processTranslate(ocrres)
}
