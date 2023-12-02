package main

import (
	"embed"
	_ "embed"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/screenshot"
	"handy-translate/translate_service"
	"log"
	"log/slog"
	"strings"

	"github.com/go-vgo/robotgo"
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

	win3 := screenshot.NewWindow(app)

	win2.On(events.Common.WindowClosing, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowClosing win2")
		win2.Hide()
	})

	win1.On(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		win1.Hide()
	})

	win3.On(events.Common.WindowClosing, func(e *application.WindowEvent) {
		app.Logger.Info("[Event] Window WindowClosing win2")
		win3.Hide()
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

	// 系统托盘
	systemTray := app.NewSystemTray()
	systemTray.SetIcon(iconlogo)
	myMenu := app.NewMenu()

	myMenu.Add("Translate").OnClick(func(ctx *application.Context) {
		win2.Show()
	})

	myMenu.Add("截图").OnClick(func(ctx *application.Context) {
		win3.SetAlwaysOnTop(true).Fullscreen().Show()
	})

	myMenu.Add("退出").OnClick(func(ctx *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(myMenu)
	systemTray.OnClick(func() {
		win1.Show()
	})

	windowMap["index"] = win1
	windowMap["translate"] = win2
	windowMap["screenshot"] = win3

	// 初始化文件和鼠标事件
	config.Init(projectName)
	go processHook()

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

// 监听处理鼠标事件
func processHook() {
	go hook.DafaultHook(windowMap) // 使用robotgo处理

	for {
		select {
		case <-hook.HookChan:
			queryText, _ := robotgo.ReadAll()
			sendQueryText(queryText)
			if queryText != translate_service.GetQueryText() {
				app.Logger.Info("GetQueryText",
					slog.String("queryText", queryText),
					slog.String("fromLang", fromLang),
					slog.String("toLang", toLang))

				processTranslate(queryText)
			}
		}
	}
}

// 翻译处理
func processTranslate(queryText string) {
	transalteWay := translate_service.GetTransalteWay(config.Data.TranslateWay)
	result, err := transalteWay.PostQuery(queryText, fromLang, toLang)
	if err != nil {
		slog.Error("PostQuery", err)
	}

	app.Logger.Info("Transalte",
		slog.Any("result", result),
		slog.Any("transalteWay", transalteWay.GetName()))

	transalteRes := strings.Join(result, "\n")

	// 发送结果至前端
	sendDataToJS(queryText, transalteRes, "")
}
