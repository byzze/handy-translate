package main

import (
	"embed"
	_ "embed"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate_service"
	"handy-translate/window/screenshot"
	"handy-translate/window/toolbar"
	"handy-translate/window/translate"
	"log"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v3/pkg/application"
	"golang.org/x/sys/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed frontend/public/appicon.png
var iconlogo []byte

var app *application.App

var appInfo = &App{}

var fromLang, toLang = "auto", "zh"

var projectName = "handy-translate"

func main() {
	app = application.New(application.Options{
		Name: projectName,
		Bind: []any{
			&App{},
		},
		Icon: iconlogo,
		Assets: application.AssetOptions{
			FS: assets,
		},
	})

	if err := mutexProcess(); err != nil {
		slog.Error("mutexProcess", slog.Any("err", err))
		return
	}

	toolbar.NewWindow(app)

	translate.NewWindow(app)

	screenshot.NewWindow(app)

	app.Events.On("translateLang", func(event *application.WailsEvent) {
		app.Logger.Info("translateType", slog.Any("event", event))

		valueType := reflect.TypeOf(event.Data)
		fmt.Println("Type:", valueType)

		if optionalData, ok := event.Data.([]interface{}); ok {
			fromLang = fmt.Sprintf("%v", optionalData[0])
			toLang = fmt.Sprintf("%v", optionalData[1])
			app.Logger.Info("translateLang",
				slog.String("fromLang", fromLang),
				slog.String("toLang", toLang))
		}
	})

	// 系统托盘
	systemTray := app.NewSystemTray()
	myMenu := app.NewMenu()

	myMenu.Add("翻译").OnClick(func(ctx *application.Context) {
		translate.Window.Center()
		translate.Window.Show()
	})

	myMenu.Add("截图").OnClick(func(ctx *application.Context) {
		screenshot.ScreenshotFullScreen()
		base64Image := screenshot.ScreenshotFullScreen()
		app.Events.Emit(&application.WailsEvent{Name: "screenshotBase64", Data: base64Image})
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
		log.Fatal(err)
	}
}

func sendDataToJS(query, result, explains string) {
	sendQueryText(query)
	sendResult(result, explains)
}

func sendQueryText(queryText string) {
	app.Events.Emit(&application.WailsEvent{Name: "query", Data: queryText})
}

func sendResult(result, explains string) {
	app.Events.Emit(&application.WailsEvent{Name: "result", Data: result})
	app.Events.Emit(&application.WailsEvent{Name: "explains", Data: explains})
}

// 监听处理鼠标事件
func processHook() {
	go hook.DafaultHook(app) // 使用robotgo处理

	for {
		select {
		case <-hook.HookChan:
			queryText, _ := robotgo.ReadAll()

			app.Logger.Info("GetQueryText",
				slog.String("queryText", queryText),
				slog.String("fromLang", fromLang),
				slog.String("toLang", toLang))

			if queryText != translate_service.GetQueryText() && queryText != "" {
				translate_service.SetQueryText(queryText)
				translateRes := processTranslate(queryText)
				// 发送结果至前端
				sendDataToJS(queryText, translateRes, "")
				continue
			}
			processToolbarShow()
		}
	}
}

// 翻译处理
func processTranslate(queryText string) string {
	translateWay := translate_service.GetTransalteWay(config.Data.TranslateWay)
	result, err := translateWay.PostQuery(queryText, fromLang, toLang)
	if err != nil {
		slog.Error("PostQuery", err)
	}

	app.Logger.Info("Transalte",
		slog.Any("result", result),
		slog.Any("translateWay", translateWay.GetName()))

	translateRes := strings.Join(result, "\n")

	return translateRes
}

func mutexProcess() error {
	// 检查是否已经存在另一个实例
	if runtime.GOOS == "windows" {
		// Windows 下使用命名管道
		mutexName := "Global\\" + projectName
		mutexNameUint, _ := syscall.UTF16PtrFromString(mutexName)
		mutexHandle, err := windows.CreateMutex(nil, false, mutexNameUint)
		if err != nil {
			var ch = make(chan int, 1)
			defer close(ch)

			go func() {
				app.Run()
			}()

			time.Sleep(time.Second * 2)

			go func() {
				dialog := application.ErrorDialog()
				dialog.SetTitle(projectName)
				dialog.SetMessage("启动失败, 另一个进程实例正在运行!!!")
				dialog.Show()
				ch <- 1
			}()

			<-ch
			app.Quit()
			return err
		}
		if mutexHandle != 0 {
			defer windows.CloseHandle(mutexHandle)
		}
	} else {
		// 其他平台使用文件锁, 暂未验证
		mutexFileName := "/tmp/" + projectName + ".lock"
		_, err := os.Create(mutexFileName)
		if err != nil {
			return err
		}
	}

	// 你的应用程序逻辑开始
	fmt.Println("Starting your desktop application...")
	return nil
}
