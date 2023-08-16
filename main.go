package main

import (
	"embed"
	"handy-translate/config"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/windows/icon.ico
var windowsIcon []byte

//go:embed build/appicon.png
var darwinIcon []byte

var appicon []byte

func init() {
	config.Init()
	// 根据操作系统选择图标文件的路径
	switch os := runtime.GOOS; os {
	case "windows":
		appicon = windowsIcon
	case "darwin":
		appicon = darwinIcon
	}
}

func main() {
	app := NewApp()
	// system tray 系统托盘
	onReady := func() {
		systray.SetIcon(appicon)
		systray.SetTitle(config.Data.Appname)
		systray.SetTooltip(config.Data.Appname + "便捷翻译工具")
		mQuitOrig := systray.AddMenuItem("退出", "退出翻译工具")
		go func() {
			logrus.Info("Requesting quit waiting")
			<-mQuitOrig.ClickedCh
			logrus.Info("Requesting quit")
			systray.Quit()
			defer app.Quit()
			logrus.Info("Finished quitting")
		}()
		// Sets the icon of a menu item. Only available on Mac and Windows.
		mQuitOrig.SetIcon(appicon)
		// Create an instance of the app structure
	}

	go systray.Run(onReady, func() { logrus.Info("app quit") })

	// Create application with options
	err := wails.Run(&options.App{
		Title:  config.Data.Appname,
		Width:  330,
		Height: 410,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// Menu: AppMenu, // reference the menu above
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:         app.startup,
		OnDomReady:        app.onDomReady,
		HideWindowOnClose: true,
		Frameless:         true,
		StartHidden:       true,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		logrus.Error("Error:", err.Error())
	}
}
