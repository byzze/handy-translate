package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"handy-translate/api/windows"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate_service"
	"handy-translate/utils"
	"image"
	"image/png"
	"log/slog"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// 和js绑定的go方法集合

// App is a service that greets people
type App struct {
}

// MyFetch URl
func (a *App) MyFetch(URL string, content map[string]interface{}) interface{} {
	return utils.MyFetch(URL, content)
}

// Transalte 翻译逻辑
func (a *App) Transalte(queryText, fromLang, toLang string) string {

	app.Logger.Info("Transalte",
		slog.Any("queryText", queryText),
		slog.Any("toLang", toLang),
		slog.Any("fromLang", fromLang))

	// 翻译loading
	app.Events.Emit(&application.WailsEvent{Name: "loading", Data: "true"})
	defer app.Events.Emit(&application.WailsEvent{Name: "loading", Data: "false"})

	transalteWay := translate_service.GetTransalteWay(config.Data.TranslateWay)

	result, err := transalteWay.PostQuery(queryText, fromLang, toLang)
	if err != nil {
		logrus.WithError(err).Error("PostQuery")
	}

	app.Logger.Info("Transalte",
		slog.Any("result", result),
		slog.Any("transalteWay", transalteWay.GetName()))

	transalteRes := strings.Join(result, "\n")

	sendDataToJS(queryText, transalteRes, "")
	return transalteRes
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
		logrus.Info("===WindowGetPosition===: ", x, y)
		sc, _ := w.GetScreen()
		logrus.Info("===GetScreen===: ", sc.Size.Width, sc.Size.Height)

		if y+w.Height() >= sc.Size.Height {
			gap := y + w.Height() - sc.Size.Height
			w.SetAbsolutePosition(x+10, y-gap-50)
		} else {
			w.SetAbsolutePosition(x+10, y+10)
		}
		windows.ShowForWindows("ToolBar")
	}
}

func (a *App) CaptureSelectedScreen(startX, startY, endwidth, endheight float64) (string, error) {
	x, y, width, height := int(startX), int(startY), int(endwidth), int(endheight)
	// 裁剪图片
	rect := image.Rect(x, y, width, height)
	if hook.IMG == nil {
		bounds := screenshot.GetDisplayBounds(0)
		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			// 错误处理，输出错误信息并返回
			fmt.Println("Error capturing screenshot:", err)
			return "", err
		}
		hook.IMG = img
	}
	croppedImg := hook.IMG.SubImage(rect)

	var buf bytes.Buffer
	err := png.Encode(&buf, croppedImg)
	if err != nil {
		return "", err
	}

	filename := "screenshot.png" // 保存的文件名
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	err = saveBase64Image(base64String, filename)
	if err != nil {
		logrus.Fatal("保存图片出错: ", err)
	}

	resut := "" //ExecOCR(".\\RapidOCR-json.exe", filename)

	return resut, nil
}

// ProcessHook 处理鼠标事件
func ProcessHook() {
	go hook.DafaultHook()    // 使用robotgo处理
	go windows.WindowsHook() // 完善，处理robotgo处理不完美
	for {
		select {
		case <-hook.HookChan:
			queryText, _ := robotgo.ReadAll()
			sendQueryText(queryText)
			if queryText != translate_service.GetQueryText() {
				app.Logger.Info("GetQueryText",
					slog.String("queryText", queryText),
					slog.String("queryText", fromLang),
					slog.String("queryText", toLang))
				appInfo.Transalte(queryText, fromLang, toLang)
			}
		}
	}
}
