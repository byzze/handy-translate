package main

import (
	"embed"
	"handy-translate/config"
	"handy-translate/log"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	config.Init()
}

func main() {
	file := log.Init()
	if file != nil {
		defer file.Close()
		file.Seek(0, 0) // 每次运行清空日志
		mw := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(mw)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  config.Data.Appname,
		Width:  330,
		Height: 400,
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
		println("Error:", err.Error())
	}
}
