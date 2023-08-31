package hook

import (
	"handy-translate/config"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

var HookChan = make(chan struct{}, 1)

var HookKeyboardChan = make(chan struct{}, 1)

var HookCPChan = make(chan struct{}, 1)

var curContent string

var lk sync.RWMutex

func SetCurText(value string) {
	lk.Lock()
	curContent = value
	lk.Unlock()
}

func GetCurText() string {
	lk.RLock()
	defer lk.RUnlock()
	return curContent
}

var defaulthook = func(e hook.Event) {
	if e.Button == hook.MouseMap["center"] {
		if pressLock.TryLock() {
			time.Sleep(time.Millisecond * 100)
			HookChan <- struct{}{}
			pressLock.Unlock()
		}
	}
}

var keyboardhook = func(e hook.Event) {
	if pressLock.TryLock() {
		logrus.Info(e)
		robotgo.KeyUp("ctrl")
		time.Sleep(time.Millisecond * 300)
		HookChan <- struct{}{}
		pressLock.Unlock()
	}
}

// Hook register hook event
func DafaultHook() {
	hook.Register(hook.MouseHold, []string{}, defaulthook)
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
	// evChan := hook.Start()
	// defer hook.End()
	// for ev := range evChan {
	// 	if ev.Button == hook.MouseMap["ctenter"] && ev.Kind == hook.MouseHold {
	// 		// 模拟按下 Ctrl+C
	// 		robotgo.KeyTap("c", "ctrl")
	// 		HookCenterChan <- struct{}{}
	// 	}
	// }
	hook.End()
	SetCurText("")
	if len(config.Data.Keyboard) == 0 || config.Data.Keyboard[0] == "center" {
		hook.Register(hook.MouseHold, []string{}, defaulthook)
	} else {
		hook.Register(hook.KeyHold, config.Data.Keyboard, keyboardhook)
	}

	s := hook.Start()
	<-hook.Process(s)

}
