package hook

import (
	"fmt"
	"handy-translate/config"
	"sync"

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

// Hook register hook event
func Hook(keyboard []string) {
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
	if len(config.Data.Keyboard) == 0 || config.Data.Keyboard[0] == "center" {
		hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
			if e.Button == hook.MouseMap["center"] {
				HookChan <- struct{}{}
			}
		})
	} else {
		hook.End()
		hook.Register(hook.KeyHold, keyboard, func(e hook.Event) {
			fmt.Println(e)
			HookChan <- struct{}{}
		})
	}

	s := hook.Start()
	<-hook.Process(s)

	// 这个会阻塞事件
	// centerBtn := robotgo.AddEvent("center")

	// // mouse center press
	// for {
	// 	if centerBtn {
	// 		robotgo.KeyTap("c", "ctrl")
	// 		HookCenterChan <- struct{}{}
	// 		robotgo.MilliSleep(100)
	// 	}

	// 	centerBtn = robotgo.AddEvent("center")
	// }
}
