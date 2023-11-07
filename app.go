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
	runtime.WindowShow(ctx)
	a.SendDataToJS("启动成功", "", "")
	// system tray 系统托盘
	onReady := func() {
		systray.SetIcon(appicon)
		systray.SetTitle(config.Data.Appname)
		systray.SetTooltip(config.Data.Appname + "便捷翻译工具")
		mShow := systray.AddMenuItem("显示", "显示翻译工具")
		mQuitOrig := systray.AddMenuItem("退出", "退出翻译工具")
		// Sets the icon of a menu item. Only available on Mac and Windows.
		mShow.SetIcon(appicon)
		for {
			select {
			case <-mShow.ClickedCh:
				a.Show()
			case <-mQuitOrig.ClickedCh:
				a.Quit()
			default:
				logrus.Println("检测系统托盘进程中")
			}
		}
	}
	systray.Run(onReady, func() { logrus.Info("app quit") })
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

				if queryText != hook.GetQueryText() {
					hook.SetQueryText(queryText)
					// 加载动画loading
					runtime.EventsEmit(a.ctx, "loading", "true")

					transalteWay := translate.GetTransalteWay(config.Data.TranslateWay)

					logrus.WithFields(logrus.Fields{
						"queryText":    queryText,
						"transalteWay": transalteWay.GetName(),
					}).Info("Transalte")

					curName := transalteWay.GetName()
					// 使用 strings.Replace 替换 \r 和 \n 为空格
					queryTextTmp := strings.ReplaceAll(queryText, "\r", "")
					queryTextTmp = strings.ReplaceAll(queryTextTmp, "\n", "")

					result, err := transalteWay.PostQuery(queryTextTmp)
					if err != nil {
						logrus.WithError(err).Error("PostQuery")
						continue
					}

					logrus.WithFields(logrus.Fields{
						"result": result,
					}).Info("Transalte")

					if len(result) >= 2 && curName == youdao.Way {
						a.SendDataToJS(queryText, result[0], result[1])
						continue
					}
					transalteRes := strings.Join(result, ",")
					a.SendDataToJS(queryText, transalteRes, "")
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

func (a *App) Quit() {
	runtime.Quit(a.ctx)
	systray.Quit()
}

func (a *App) Show() {
	runtime.WindowShow(a.ctx)
}
