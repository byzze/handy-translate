package main

import (
	"embed"
	_ "embed"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate_service"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed frontend/dist
var assets embed.FS

var app *application.App

var appInfo = &App{}

var windowMap = make(map[string]*application.WebviewWindow)

var W1 *application.WebviewWindow

func main() {
	var projectName = "handy-translate"
	app = application.New(application.Options{
		Name: projectName,
		Bind: []any{
			appInfo,
		},
		Assets: application.AssetOptions{
			FS: assets,
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create window
	W1 = app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:  "ToolBar",
		Width:  300,
		Height: 500,
		ShouldClose: func(window *application.WebviewWindow) bool {
			app.Quit()
			return true
		},
		URL: "index.html",
	})

	win2 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Translate",
		Width:     500,
		Height:    500,
		Frameless: true,
		Hidden:    true,
		URL:       "translate.html",
	})

	win2.On(events.Common.WindowClosing, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowClosing! home1")
		win2.Hide()
	})

	win2.On(events.Common.WindowFocus, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowFocus! home2")
	})

	app.Events.On("translateLang", func(event *application.WailsEvent) {
		logrus.WithField("event", event).Info("translateType")
		if optionalData, ok := event.Data.([]string); ok {
			fromLang = fmt.Sprintf("%v", optionalData[0])
			toLang = fmt.Sprintf("%v", optionalData[1])
			logrus.Info(fromLang, toLang)
		}
	})

	windowMap["index"] = W1
	windowMap["translate"] = win2

	app.On(events.Common.ApplicationStarted, func(event *application.Event) {
		config.Init(projectName)
		go ProcessHook()
		go hook.DafaultHook()
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func ProcessHook() {
	time.Sleep(time.Second * 3)
	for {
		select {
		case <-hook.HookChan:
			robotgo.KeyTap("c", "ctrl")
			// // x, y := robotgo.GetMousePos()
			// // logrus.Info("===WindowGetPosition===: ", x, y)
			// // runtime.WindowSetPosition(ctx, x, y)
			// // logrus.Info("HookChan Process")
			// // runtime.WindowShow(ctx)
			// // windowX, windowY := runtime.WindowGetSize(ctx)
			// // x, y := robotgo.GetMousePos()
			// // x, y = x+10, y-10

			// // runtime.WindowFullscreen(ctx)
			// // runtime.EventsEmit(ctx, "appLabel", "translate")
			queryText, _ := robotgo.ReadAll()
			logrus.Info("===ClipboardGetText===: ", queryText)
			x, y := robotgo.Location()
			logrus.Info("===WindowGetPosition===: ", x, y)
			W1.SetAbsolutePosition(x+10, y+10)
			sendQueryText(queryText)
			if queryText != translate_service.GetQueryText() {
				fmt.Println("GetQueryText", fromLang, toLang)
				appInfo.Transalte(queryText, fromLang, toLang)
			}
			W1.Show()
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
}
