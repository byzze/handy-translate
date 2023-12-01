package hook

import (
	"bytes"
	"encoding/base64"
	"handy-translate/config"
	"image"
	"image/png"

	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

// HookChan
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

// 将图像编码为Base64字符串
func encodeImageToBase64(img image.Image) string {
	// 创建一个缓冲区用于保存Base64编码的数据
	var imgBytes []byte
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		panic(err)
	}

	imgBytes = buf.Bytes()

	// 使用base64编码图像数据
	base64Image := base64.StdEncoding.EncodeToString(imgBytes)

	return base64Image
}

var lastKeyPressTime time.Time

var lastMouseTime time.Time

var IMG *image.RGBA

// DafaultHook register hook event
func DafaultHook() {
	// hook.Register(hook.KeyDown, []string{"ctrl", "c", "c"}, func(e hook.Event) {
	// 	logrus.Info(e)
	// 	if pressLock.TryLock() {
	// 		lastKeyPressTime = time.Now()
	// 	} else {
	// 		elapsed := time.Since(lastKeyPressTime)
	// 		// Check if the time elapsed is greater than 500 milliseconds
	// 		if elapsed.Milliseconds() < 800 {
	// 			// HookChan <- struct{}{}
	// 		}
	// 		pressLock.Unlock()
	// 	}
	// })

	// hook.Register(hook.KeyDown, []string{"f", "ctrl", "shift"}, func(e hook.Event) {
	// 	logrus.Info(e)
	// 	bounds := screenshot.GetDisplayBounds(0)
	// 	img, err := screenshot.CaptureRect(bounds)

	// 	if err != nil {
	// 		// 错误处理，输出错误信息并返回
	// 		fmt.Println("Error capturing screenshot:", err)
	// 		return
	// 	}
	// 	IMG = img

	// 	base64Image := encodeImageToBase64(img)
	// 	if base64Image == "" {
	// 		// 错误处理，未能生成Base64图像，返回
	// 		fmt.Println("Error encoding image to Base64")
	// 		return
	// 	}

	// 	runtime.EventsEmit(ctx, "appLabel", "screenshot")

	// 	runtime.EventsEmit(ctx, "screenshot", base64Image)
	// })

	hook.Register(hook.MouseDown, []string{}, defaulthook)

	s := hook.Start()
	<-hook.Process(s)
}

var pressLock sync.RWMutex

// Hook register hook event 用于配置快捷键 TODO
func Hook() {
	logrus.Info("--- Please wait hook starting ---")
	hook.End()
	if len(config.Data.Keyboard) == 0 || config.Data.Keyboard[0] == "center" {
		hook.Register(hook.MouseDown, []string{}, defaulthook)
	} else {
		hook.Register(hook.KeyDown, config.Data.Keyboard, keyboardhook)
	}

	s := hook.Start()
	<-hook.Process(s)
}
