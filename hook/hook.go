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

// Hook register hook event
func DafaultHook() {
	hook.Register(hook.MouseDown, []string{}, defaulthook)
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
