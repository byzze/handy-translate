package main

import (
	"context"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate"
	"handy-translate/translate/youdao"
	"handy-translate/utils"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) MyFetch(URL string, content map[string]interface{}) interface{} {
	return utils.MyFetch(URL, content)
}

func (a *App) sendQueryText(queryText string) {
	runtime.EventsEmit(a.ctx, "query", queryText)
}

func (a *App) sendResult(result, explian string) {
	runtime.EventsEmit(a.ctx, "result", result)
	runtime.EventsEmit(a.ctx, "explian", explian)
}

func (a *App) SendDataToJS(query, result, explian string) {
	logrus.WithFields(logrus.Fields{
		"query":   query,
		"result":  result,
		"explian": explian,
	}).Info("SendDataToJS", query, result, explian)

	a.sendQueryText(query)
	a.sendResult(result, explian)

}

func (a *App) onDomReady(ctx context.Context) {
	a.sendQueryText("启动成功\n唤醒应用: 鼠标中键或者Ctrl+c+c\nOCR截图翻译: Ctrl+Shift+F ")
}

var fromLang, toLang = "auto", "zh"

func eventFunc(ctx context.Context) {
	runtime.EventsOn(ctx, "translateType", func(optionalData ...interface{}) {
		logrus.WithField("optionalData", optionalData).Info("translateType")
		if len(optionalData) >= 2 {
			fromLang = fmt.Sprintf("%v", optionalData[0])
			toLang = fmt.Sprintf("%v", optionalData[1])
		}
	})

	runtime.EventsOn(ctx, "screenshotCapture", func(optionalData ...interface{}) {
		logrus.WithField("screenshotCapture", optionalData).Info("translateType")
		if len(optionalData) >= 1 {
			base64String := fmt.Sprintf("%v", optionalData[0])
			base64String = strings.TrimPrefix(base64String, "data:image/png;base64,")
			logrus.WithField("base64String", base64String).Info("translateType")
			filename := "screenshot.png" // 保存的文件名

			err := saveBase64Image(base64String, filename)
			if err != nil {
				logrus.Fatal("保存图片出错: ", err)
			}
			logrus.Println("图片保存成功")
			resut := ExecOCR(".\\RapidOCR-json.exe", filename)
			runtime.EventsEmit(ctx, "query", resut)
		}
	})
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	config.Init(ctx)

	go hook.DafaultHook(ctx)
	go hook.WindowsHook()

	eventFunc(ctx)
	// scList, _ := runtime.ScreenGetAll(ctx)

	// var screenX, screenY int
	// for _, v := range scList {
	// 	if v.IsCurrent {
	// 		screenX = v.Width
	// 		screenY = v.Height
	// 	}
	// }

	runtime.WindowCenter(ctx)
	go func() {
		for {
			select {
			case <-hook.HookChan:
				logrus.Info("HookChan Process")
				// windowX, windowY := runtime.WindowGetSize(ctx)
				// x, y := robotgo.GetMousePos()
				// x, y = x+10, y-10

				// runtime.WindowFullscreen(ctx)
				runtime.EventsEmit(ctx, "ocrShow", false)
				queryText, _ := runtime.ClipboardGetText(a.ctx)

				a.sendQueryText(queryText)
				if queryText != hook.GetQueryText() {
					fmt.Println("GetQueryText", fromLang, toLang)
					a.Transalte(queryText, fromLang, toLang)
				}

				// TODO 弹出窗口根据鼠标位置变动
				// fmt.Println("or:", x, y, screenX, screenY, windowX, windowY)
				// if y+windowY+20 >= screenY {
				// 	y = screenY - windowY - 20
				// }

				// if x+windowX >= screenX {
				// 	x = screenX - windowX
				// }
				// fmt.Println("new:", x, y, screenX, screenY, windowX, windowY)
				// runtime.WindowSetPosition(ctx, x, y)
			}
		}
	}()
}

// Greet returns a greeting for the given name
func (a *App) GetKeyBoard() []string {
	if len(config.Data.Keyboard) == 0 {
		config.Data.Keyboard = make([]string, 3)
	}
	return config.Data.Keyboard
}

func (a *App) SetKeyBoard(ctrl, shift, key string) {
	config.Data.Keyboard = []string{ctrl, shift, key}
	logrus.Info(config.Data.Keyboard)
	config.Save()
	go hook.Hook()
}

func (a *App) GetTransalteMap() string {
	var translateList = config.Data.Translate
	bTranslate, err := json.Marshal(translateList)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}
	return string(bTranslate)
}

func (a *App) SetTransalteWay(translateWay string) {
	fmt.Println(translateWay)
	config.Data.TranslateWay = translateWay
	hook.SetQueryText("")
	config.Save()
	logrus.WithField("config.Data.Translate", config.Data.Translate).Info("SetTransalteList")
}

func (a *App) GetTransalteWay() string {
	return config.Data.TranslateWay
}

func (a *App) Transalte(queryText, fromLang, toLang string) {
	hook.SetQueryText(queryText)
	// 加载动画loading
	runtime.EventsEmit(a.ctx, "loading", "true")
	defer runtime.EventsEmit(a.ctx, "loading", "false")

	transalteWay := translate.GetTransalteWay(config.Data.TranslateWay)

	logrus.WithFields(logrus.Fields{
		"queryText":    queryText,
		"transalteWay": transalteWay.GetName(),
		"fromLang":     fromLang,
		"toLang":       toLang,
	}).Info("Transalte")

	curName := transalteWay.GetName()
	// 使用 strings.Replace 替换 \r 和 \n 为空格

	result, err := transalteWay.PostQuery(queryText, fromLang, toLang)
	if err != nil {
		logrus.WithError(err).Error("PostQuery")
	}

	logrus.WithFields(logrus.Fields{
		"result": result,
	}).Info("Transalte")

	if len(result) >= 2 && curName == youdao.Way {
		a.SendDataToJS(queryText, result[0], result[1])
	}

	transalteRes := strings.Join(result, ",")
	a.SendDataToJS(queryText, transalteRes, "")

}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
	// systray.Quit()
}

func (a *App) Show() {
	runtime.WindowShow(a.ctx)
}
