package main

import (
	"embed"
	_ "embed"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate_service"
	"log"

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

	app.On(events.Common.ApplicationStarted, func(event *application.Event) {
		config.Init(projectName)
		go ProcessHook()
		go hook.DafaultHook()
	})

	// Create window
	win1 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "ToolBar",
		URL:   "index.html",
	})

	win2 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:  "Translate",
		Width:  500,
		Height: 500,
		// Frameless: true,
		// Hidden: true,
		URL: "translate.html",
	})

	win1.On(events.Common.WindowFocus, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowFocus! home1")
	})

	win2.On(events.Common.WindowFocus, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowFocus! home2")
	})

	app.Events.On("translateType", func(event *application.WailsEvent) {
		logrus.WithField("event", event).Info("translateType")
		// if optionalData, ok := event.Data([]string); ok {
		// 	fromLang = fmt.Sprintf("%v", optionalData[0])
		// 	toLang = fmt.Sprintf("%v", optionalData[1])
		// }
	})

	windowMap["index"] = win1
	windowMap["translate"] = win2

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func ProcessHook() {
	for {
		select {
		case <-hook.HookChan:
			// robotgo.KeyTap("c", "ctrl")
			// x, y := robotgo.GetMousePos()
			// logrus.Info("===WindowGetPosition===: ", x, y)
			// runtime.WindowSetPosition(ctx, x, y)
			// logrus.Info("HookChan Process")
			// runtime.WindowShow(ctx)
			// windowX, windowY := runtime.WindowGetSize(ctx)
			// x, y := robotgo.GetMousePos()
			// x, y = x+10, y-10

			// runtime.WindowFullscreen(ctx)
			// runtime.EventsEmit(ctx, "appLabel", "translate")
			queryText, _ := robotgo.ReadAll()
			logrus.Info("===ClipboardGetText===: ", queryText)
			x, y := robotgo.Location()
			logrus.Info("===WindowGetPosition===: ", x, y)
			windowMap["index"].SetAbsolutePosition(x+10, y+10)
			sendQueryText(queryText)
			if queryText != translate_service.GetQueryText() {
				fmt.Println("GetQueryText", fromLang, toLang)
				appInfo.Transalte(queryText, fromLang, toLang)
			}
			windowMap["index"].Show()
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
