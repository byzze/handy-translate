package hook

import (
	"bytes"
	"context"
	"encoding/base64"
	"handy-translate/config"
	"image"
	"image/png"

	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var HookChan = make(chan struct{}, 1)

var queryText string

var lk sync.RWMutex

func SetQueryText(value string) {
	lk.Lock()
	queryText = value
	lk.Unlock()
}

func GetQueryText() string {
	lk.RLock()
	defer lk.RUnlock()
	return queryText
}

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

var OcrShow bool

// Hook register hook event
func DafaultHook(ctx context.Context) {
	hook.Register(hook.MouseDown, []string{}, defaulthook)
	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "1"}, func(e hook.Event) {
		if pressLock.TryLock() {
			runtime.EventsEmit(ctx, "ocrShow", true)
			logrus.Info(e)
			i := 0
			bounds := screenshot.GetDisplayBounds(i)

			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				panic(err)
			}

			base64Image := encodeImageToBase64(img)
			runtime.EventsEmit(ctx, "screenshot", base64Image)
			runtime.EventsOff(ctx, "screenshot")
			time.Sleep(time.Millisecond * 100)
			pressLock.Unlock()
		}
	})
	s := hook.Start()
	<-hook.Process(s)
	// 这个会阻塞事件
	/* centerBtn := robotgo.AddEvent("center")
	// mouse center press
	for {
		if centerBtn {
			if pressLock.TryLock() {
				HookChan <- struct{}{}
				robotgo.MilliSleep(100)
				pressLock.Unlock()
			}
		}
		centerBtn = robotgo.AddEvent("center")
	}
	*/
}

var pressLock sync.RWMutex

// Hook register hook event
func Hook() {
	logrus.Info("--- Please wait hook starting ---")
	hook.End()
	SetQueryText("")
	if len(config.Data.Keyboard) == 0 || config.Data.Keyboard[0] == "center" {
		hook.Register(hook.MouseDown, []string{}, defaulthook)
	} else {
		hook.Register(hook.KeyDown, config.Data.Keyboard, keyboardhook)
	}

	s := hook.Start()
	<-hook.Process(s)
}
