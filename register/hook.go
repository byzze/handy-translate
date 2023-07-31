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
	// 监听鼠标事件
	centerBtn := robotgo.AddEvent("center")

	// 等待鼠标中键按下
	for {
		if centerBtn {
			// 按下鼠标中键，执行复制文本到剪贴板的操作
			robotgo.MilliSleep(100)
			handleData()
			HookCenterChan <- struct{}{}
			// 等待鼠标中键释放
			robotgo.MilliSleep(100)
			centerBtn = robotgo.AddEvent("center")
		}
	}
}

func handleData() {
	// 读取原来的内容
	oldContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData oldContent ReadAll")
	}
	// 模拟按下 Ctrl 键
	robotgo.KeyTap("c", "ctrl")
	tmpContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData tmpContent ReadAll")
	}
	SetQueryContent(tmpContent)
	// 将原来的数据写回去，防止污染剪贴板
	if err := robotgo.WriteAll(oldContent); err != nil {
		logrus.WithError(err).Error("handleData WriteAll")
	}
	logrus.WithField("tmpcontent", tmpContent).Info("handleData finsh")
}
