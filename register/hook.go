package register

import (
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
func Hook() {
	logrus.Info("--- Please wait hook starting ---")
	centerBtn := robotgo.AddEvent("center")

	// mouse center press
	for {
		if centerBtn {
			robotgo.MilliSleep(100)
			handleData()
			HookCenterChan <- struct{}{}
			robotgo.MilliSleep(100)
			centerBtn = robotgo.AddEvent("center")
		}
	}
}

func handleData() {
	// read origin data
	oldContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData oldContent ReadAll")
	}
	// press Ctrl + C
	robotgo.KeyTap("c", "ctrl")
	tmpContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData tmpContent ReadAll")
	}
	SetQueryContent(tmpContent)
	// write origin data
	if err := robotgo.WriteAll(oldContent); err != nil {
		logrus.WithError(err).Error("handleData WriteAll")
	}

	logrus.WithField("tmpcontent", tmpContent).Info("handleData finsh")
}
