package main

import (
	"embed"
	_ "embed"
	"fmt"
	"handy-translate/config"
	"log"
	"log/slog"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed frontend/public/wails.png
var iconlogo []byte

var app *application.App

var appInfo = &App{}

var windowMap = make(map[string]*application.WebviewWindow)

var fromLang, toLang = "auto", "zh"

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

	win1 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:         "ToolBar",
		Width:         300,
		Height:        55,
		AlwaysOnTop:   false,
		Hidden:        true,
		DisableResize: false,
		Frameless:     true,
		Centered:      true,
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
		URL:       "translate.html",
	})

	win3 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		// HTML: `
		// <div style='width: 100%; height: 100%; border: 3px solid red; background-color: "0000";'>
		// 	<div>test</div>
		// </div>`,
		// X:    rand.Intn(1000),
		// Y:    rand.Intn(800),
		// IgnoreMouseEvents: true,
		Centered: true,
		// Frameless:      true,
		BackgroundType: application.BackgroundTypeTransparent,
		// Mac: application.MacWindow{
		// 	InvisibleTitleBarHeight: 50,
		// },
		URL: "screenshot.html",
	})
	win3.Fullscreen()
	win2.On(events.Common.WindowClosing, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowClosing win2")
		win2.Hide()
	})

	win1.On(events.Common.WindowLostFocus, func(event *application.WindowEvent) {
		logrus.Info("[Event] Window WindowLostFocus win1")
		win1.Hide()
	})

	app.Events.On("translateLang", func(event *application.WailsEvent) {
		app.Logger.Info("translateType", slog.Any("event", event))

		if optionalData, ok := event.Data.([]string); ok {
			fromLang = fmt.Sprintf("%v", optionalData[0])
			toLang = fmt.Sprintf("%v", optionalData[1])
			app.Logger.Info("translateLang",
				slog.String("fromLang", fromLang),
				slog.String("toLang", toLang))
		}
	})

	windowMap["index"] = win1
	windowMap["translate"] = win2
	windowMap["screenshot"] = win3

	// 系统托盘
	systemTray := app.NewSystemTray()
	systemTray.SetIcon(iconlogo)
	myMenu := app.NewMenu()

	myMenu.Add("显示").OnClick(func(ctx *application.Context) {
		win3.Show()
	})

	myMenu.Add("退出").OnClick(func(ctx *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(myMenu)
	systemTray.OnClick(func() {
		win3.Show()
	})

	// 初始化文件和鼠标事件
	config.Init(projectName)
	// go ProcessHook()

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func sendDataToJS(query, result, explian string) {
	sendQueryText(query)
	sendResult(result, explian)
}

func sendQueryText(queryText string) {
	app.Events.Emit(&application.WailsEvent{Name: "query", Data: queryText})
}

func sendResult(result, explian string) {
	app.Events.Emit(&application.WailsEvent{Name: "result", Data: result})
	app.Events.Emit(&application.WailsEvent{Name: "explian", Data: explian})
}
