package hook

import (
	"context"
	"sync"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
)

var HookCenterChan = make(chan struct{}, 1)

var queryContent string
var curContent string
var err error

var lk sync.RWMutex
var datalk sync.RWMutex

// SetQueryContent lock data
func SetQueryContent(value string) {
	lk.Lock()
	queryContent = value
	lk.Unlock()
}

func GetQueryText() string {
	lk.RLock()
	defer lk.RUnlock()
	return queryContent
}

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
			handleData()
			HookCenterChan <- struct{}{}
			robotgo.MilliSleep(100)
		}
		centerBtn = robotgo.AddEvent("center")
	}
}

func handleData() {
	// press Ctrl + C
	robotgo.KeyTap("c", "ctrl")
	robotgo.MilliSleep(50)
	tmpContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData tmpContent ReadAll")
	}
	logrus.WithField("tmpcontent", tmpContent).Info("handleData finsh")
	SetQueryContent(tmpContent)
}
