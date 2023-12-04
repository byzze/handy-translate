package hook

import (
	"handy-translate/config"
	"handy-translate/os_api/windows"
	"handy-translate/screenshot"
	"log/slog"
	"runtime"

	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// HookChan channle
var HookChan = make(chan struct{}, 1)

var defaulthook = func(e hook.Event) {
	if e.Button == hook.MouseMap["center"] {
		HookChan <- struct{}{}
	}
}

var keyboardhook = func(e hook.Event) {
	if pressLock.TryLock() {
		logrus.Info(e)
		robotgo.KeyTap("c", "ctrl")
		time.Sleep(time.Millisecond * 100)
		HookChan <- struct{}{}
		pressLock.Unlock()
	}
}

var lastKeyPressTime time.Time

var lastMouseTime time.Time

// DafaultHook register hook event
func DafaultHook(app *application.App) {
	if runtime.GOOS == "windows" {
		go windows.WindowsHook() // 完善，robotgo处理的不完美
	}

	hook.Register(hook.KeyDown, []string{"c", "c", "ctrl"}, func(e hook.Event) {
		slog.Info("ctrl+c+c", e)
		if pressLock.TryLock() {
			lastKeyPressTime = time.Now()
		} else {
			pressLock.Unlock()
			elapsed := time.Since(lastKeyPressTime)
			// Check if the time elapsed is greater than 500 milliseconds
			if elapsed.Milliseconds() < 800 {
				HookChan <- struct{}{}
			}

		}
	})

	screenshotKey := config.Data.Keyboards["screenshot"]
	hook.Register(hook.KeyDown, screenshotKey, func(e hook.Event) {
		slog.Info("screenshot", e)
		base64Image := screenshot.ScreenshotFullScreen()
		app.Events.Emit(&application.WailsEvent{Name: "screenshotBase64", Data: base64Image})

	})

	// default mid mouse
	hook.Register(hook.MouseDown, []string{}, defaulthook)

	s := hook.Start()
	<-hook.Process(s)
}

var pressLock sync.RWMutex

// ToolBarHook register hook event 用于配置快捷键 TODO
func ToolBarHook() {
	logrus.Info("--- Please wait hook starting ---")
	hook.End()
	if len(config.Data.Keyboards) == 0 || config.Data.Keyboards["center"][0] == "center" {
		hook.Register(hook.MouseDown, []string{}, defaulthook)
	} else {
		hook.Register(hook.KeyDown, config.Data.Keyboards["toolBar"], keyboardhook)
	}

	s := hook.Start()
	<-hook.Process(s)
}
