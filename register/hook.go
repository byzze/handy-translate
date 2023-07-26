package register

import (
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

var HookChan = make(chan hook.Event, 1)

var queryContent string
var curContent string
var err error

var lk sync.RWMutex
var datalk sync.RWMutex

// 操作数据加锁
func SetCurContent(value string) {
	lk.Lock()
	curContent = value
	lk.Unlock()
}

// 操作数据加锁
func SetQueryContent(value string) {
	lk.Lock()
	queryContent = value
	lk.Unlock()
}

// 读取数据
func GetQueryContent() string {
	lk.RLock()
	defer lk.RUnlock()
	return queryContent
}

// 读取数据
func GetCurContent() string {
	lk.RLock()
	defer lk.RUnlock()
	return curContent
}

// Hook register hook event
func Hook() {
	logrus.Info("--- Please wait hook starting ---")
	evChan := hook.Start()
	defer hook.End()
	var preKind uint8
	for ev := range evChan {
		switch ev.Kind {
		case hook.HookEnabled:
			logrus.Info("--- Please hook start success ---")
		case hook.MouseMove:
			continue
		case hook.MouseUp:
			logrus.Info("MouseUp: ", ev)
			if ev.Button == hook.MouseMap["left"] && (ev.Clicks == 2 || ev.Clicks == 3) {
				handleData("left MouseUp")
				continue
			}
		case hook.MouseDown:
			if ev.Button == hook.MouseMap["center"] {
				logrus.Info("center: ", ev)
				HookChan <- ev
			}
			if ev.Button == hook.MouseMap["left"] && preKind == hook.MouseDrag {
				handleData("left Moseup")
			}
		}
		preKind = ev.Kind
		// logrus.WithField("hook: ", ev).Info()
	}
}

func handleData(mouse string) {
	datalk.Lock()
	defer datalk.Unlock()
	logrus.Info("handleData start", mouse)
	// 读取原来的内容
	oldContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData oldContent ReadAll")
		return
	}
	// 模拟按下 Ctrl 键
	robotgo.KeyTap("c", "ctrl")
	tmpContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData tmpContent ReadAll")
		return
	}
	SetQueryContent(tmpContent)
	// 将原来的数据写回去，防止污染剪贴板
	if err := robotgo.WriteAll(oldContent); err != nil {
		logrus.WithError(err).Error("handleData WriteAll")
	}
	logrus.Info("handleData end")
	time.Sleep(time.Millisecond)
}
