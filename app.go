package main

import (
	"context"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate"
	"handy-translate/translate/youdao"
	"strings"

	"github.com/getlantern/systray"
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

func (a *App) SendDataToJS(query, result, explian string) {
	logrus.WithFields(logrus.Fields{
		"query":   query,
		"result":  result,
		"explian": explian,
	}).Info("SendDataToJS", query, result, explian)

	runtime.EventsEmit(a.ctx, "query", query)
	runtime.EventsEmit(a.ctx, "result", result)
	runtime.EventsEmit(a.ctx, "explian", explian)
}

// test data
func (a *App) onDomReady(ctx context.Context) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   config.Data.Appname + " 启动成功",
		Message: "选中英文或中文按压鼠标中键即可弹出界面显示翻译结果！！！",
		// DefaultButton: "No",
	})
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	config.Init(ctx)

	go hook.DafaultHook()
	go hook.WindowsHook()
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
				runtime.WindowShow(ctx)
				queryText, _ := runtime.ClipboardGetText(a.ctx)

				if queryText != hook.GetCurText() {
					hook.SetCurText(queryText)
					// 加载动画loading
					runtime.EventsEmit(a.ctx, "loading", "true")
					wayList := translate.GetTransalteWay()
					var result []string
					var err error
					var curName string
					for i := range wayList {
						obj := wayList[i]
						logrus.WithFields(logrus.Fields{
							"queryText":    queryText,
							"transalteWay": obj.GetName(),
						}).Info("Transalte")
						curName = obj.GetName()
						// 使用 strings.Replace 替换 \r 和 \n 为空格
						queryText = strings.ReplaceAll(queryText, "\r", " ")
						queryText = strings.ReplaceAll(queryText, "\n", " ")

						result, err = obj.PostQuery(queryText)
						if err != nil {
							logrus.WithError(err).Error("PostQuery")
							continue
						}

						logrus.WithFields(logrus.Fields{
							"result": result,
						}).Info("Transalte")

						if len(result) == 0 {
							continue
						}
						break
					}
					if len(result) >= 2 && curName == youdao.Way {
						a.SendDataToJS(queryText, result[0], result[1])
						continue
					}
					a.SendDataToJS(queryText, strings.Join(result, ","), "")
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
	a.ctx = ctx
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

func (a *App) GetTransalteList() string {
	var translateList = config.Data.Translate
	bTranslate, err := json.Marshal(translateList)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}
	return string(bTranslate)
}

func (a *App) SetTransalteList(result string) {
	fmt.Println(result)
	err := json.Unmarshal([]byte(result), &config.Data.Translate)
	if err != nil {
		logrus.WithError(err).Error("SetTransalteList")
	}
	hook.SetCurText("")
	config.Save()
	logrus.WithField("config.Data.Translate", config.Data.Translate).Info("SetTransalteList")
}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
	systray.Quit()
}

func (a *App) Show() {
	runtime.WindowShow(a.ctx)
}
