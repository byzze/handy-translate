package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"handy-translate/config"
	"handy-translate/os_api/windows"
	"handy-translate/screenshot"
	"handy-translate/translate_service"
	"handy-translate/utils"
	"image/png"
	"log/slog"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
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

	/* 	// 翻译loading
	   	app.Events.Emit(&application.WailsEvent{Name: "loading", Data: "true"})
	   	defer app.Events.Emit(&application.WailsEvent{Name: "loading", Data: "false"}) */

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
	translate_service.SetQueryText("")
	config.Save()
	logrus.WithField("Translate", config.Data.Translate).Info("SetTransalteList")
}

// GetTransalteWay 获取当前翻译的服务
func (a *App) GetTransalteWay() string {
	return config.Data.TranslateWay
}

// Show 通过名字控制窗口事件
func (a *App) Show(windowName string) {
	if _, ok := windowMap[windowName]; ok {
		windowMap[windowName].Center()
		windowMap[windowName].Show()
	}
}

// Hide 通过名字控制窗口事件
func (a *App) Hide(windowName string) {
	if _, ok := windowMap[windowName]; ok {
		windowMap[windowName].Hide()
	}
}

// ToolBarShow 显示工具弹窗，控制大小，布局
func (a *App) ToolBarShow(height float64) {
	app.Logger.Info("ToolBarShow", slog.Float64("height", height))

	h := int(height)
	if h > 600 {
		h = 600
	}

	if h == 0 {
		h = 55
	}

	if w, ok := windowMap["index"]; ok {
		w.SetSize(w.Width(), h)
		x, y := robotgo.Location()
		sc, _ := w.GetScreen()

		if y+w.Height() >= sc.Size.Height {
			gap := y + w.Height() - sc.Size.Height
			w.SetAbsolutePosition(x+10, y-gap-50)
		} else {
			w.SetAbsolutePosition(x+10, y+10)
		}
		// windows.FindWindow("ToolBar").ShowForWindows() // 使用原生showwindow，wails3版本有些问题，无法正常显示
		windows.ShowForWindows("ToolBar") // 使用原生showwindow，wails3版本有些问题，无法正常显示
	}
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
