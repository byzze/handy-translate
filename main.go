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

var W1 *application.WebviewWindow

func main() {
	var projectName = "handy-translate"
	app = application.New(application.Options{
		Name: projectName,
		Bind: []any{
			&App{},
		},
		Assets: application.AssetOptions{
			FS: assets,
		},
	})

	// Create window
	W1 = app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "ToolBar",
		Width:     300,
		Height:    55,
		Centered:  true,
		Focused:   true,
		Frameless: true,
		ShouldClose: func(window *application.WebviewWindow) bool {
			app.Quit()
			return true
		},
		// Windows: application.WindowsWindow{
		// 	HiddenOnTaskbar: true,
		// },
		URL: "index.html",
	})

	win2 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Translate",
		Width:     500,
		Height:    500,
		Frameless: true,
		Hidden:    true,

		URL: "translate.html",
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

	W1.On(events.Common.WindowLostFocus, func(event *application.WindowEvent) {
		logrus.Info("WindowLostFocus")
		W1.Hide()
	})

	W1.On(events.Common.WindowShow, func(event *application.WindowEvent) {
		logrus.Info("events.Common.WindowShow")
	})

	windowMap["index"] = W1
	windowMap["translate"] = win2

	app.On(events.Common.ApplicationStarted, func(event *application.Event) {
		config.Init(projectName)
		go ProcessHook()
		go hook.DafaultHook()
		go hook.WindowsHook()
	})

	systemTray := app.NewSystemTray()
	myMenu := app.NewMenu()
	myMenu.Add("Show").OnClick(func(ctx *application.Context) {
		W1.Show()
	})

	myMenu.Add("Quit").OnClick(func(ctx *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(myMenu)
	systemTray.OnClick(func() {
		W1.Show()
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func ProcessHook() {
	for {
		select {
		case <-hook.HookChan:
			queryText, _ := robotgo.ReadAll()
			logrus.Info("===ClipboardGetText===: ", queryText)
			sendQueryText(queryText)
			if queryText != translate_service.GetQueryText() {
				fmt.Println("GetQueryText", fromLang, toLang)
				appInfo.Transalte(queryText, fromLang, toLang)
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
}
