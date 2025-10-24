package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"time"

	"handy-translate/config"
	"handy-translate/window/screenshot"
	"handy-translate/window/toolbar"
	"handy-translate/window/translate"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed frontend/public/appicon.png
var iconlogo []byte

var app *application.App

var fromLang, toLang = "auto", "zh"

var projectName = "handy-translate"

func main() {
	app = application.New(application.Options{
		Name: projectName,
		Services: []application.Service{
			application.NewService(&App{}),
		},
		Icon: iconlogo,
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "com.wails.handy-translate",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				log.Printf("Second instance launched with args: %v", data.Args)
				log.Printf("Working directory: %s", data.WorkingDir)
				log.Printf("Additional data: %v", data.AdditionalData)
			},
			// Optional: Pass additional data to second instance
			AdditionalData: map[string]string{
				"launchtime": time.Now().String(),
			},
		},
	})

	toolbar.NewWindow(app)

	translate.NewWindow(app)

	screenshot.NewWindow(app)

	app.Event.On("translateLang", func(event *application.CustomEvent) {
		app.Logger.Info("translateType", slog.Any("event", event))

		if dataSlice, ok := event.Data.([]interface{}); ok {
			if len(dataSlice) >= 2 {
				fromLang = fmt.Sprintf("%v", dataSlice[0])
				toLang = fmt.Sprintf("%v", dataSlice[1])
				app.Logger.Info("translateLang",
					slog.String("fromLang", fromLang),
					slog.String("toLang", toLang))
			}
		}
	})

	// 系统托盘
	systemTray := app.SystemTray.New()
	myMenu := app.Menu.New()

	myMenu.Add("翻译").OnClick(func(ctx *application.Context) {
		translate.Window.Center()
		translate.Window.Show()
	})

	myMenu.Add("截图").OnClick(func(ctx *application.Context) {
		screenshot.ScreenshotFullScreen()
		base64Image := screenshot.ScreenshotFullScreen()
		app.Event.Emit("screenshotBase64", base64Image)
	})

	myMenu.Add("退出").OnClick(func(ctx *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(myMenu)
	systemTray.SetIcon(iconlogo)

	systemTray.OnClick(func() {
		toolbar.Window.Show()
	})

	// 初始化文件和鼠标事件
	config.Init(projectName)
	go processHook()

	err := app.Run()
	if err != nil {
		// 报错退出程序
		panic(err)
	}
}
