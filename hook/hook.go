package hook

import (
	"context"
	"sync"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
)

var HookCenterChan = make(chan struct{}, 1)

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
func Hook(ctx context.Context) {
	logrus.Info("--- Please wait hook starting ---")

	centerBtn := robotgo.AddEvent("center")

	// mouse center press
	for {
		if centerBtn {
			robotgo.KeyTap("c", "ctrl")
			HookCenterChan <- struct{}{}
			robotgo.MilliSleep(100)
		}

		centerBtn = robotgo.AddEvent("center")
	}
}
