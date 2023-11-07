package main

import (
	"embed"
	"handy-translate/config"
	"runtime"
	"sync"

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

	// 根据操作系统选择图标文件的路径
	switch os := runtime.GOOS; os {
	case "windows":
		appicon = windowsIcon
	case "darwin":
		appicon = darwinIcon
	}
}

var sc sync.Once

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	sc.Do(func() {
		app := NewApp()

		// Create application with options
		err := wails.Run(&options.App{
			Title:  config.Data.Appname,
			Width:  460,
			Height: 410,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
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
		systray.Quit()
	})
}
