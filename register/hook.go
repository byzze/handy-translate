package register

import (
	"strings"
	"sync"
	"time"
	"translate/mywindown"
	"translate/translate"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

var queryContent string
var curContent string
var err error

var lk sync.RWMutex

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

// Hook register hook event
func Hook(queryContentLab, transalteResultLab, transalteExplainsLab *widget.Label) {
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
			if ev.Button == hook.MouseMap["left"] && (ev.Clicks == 2 || ev.Clicks == 3) {
				// handleData()
				continue
			}

			if ev.Button == hook.MouseMap["center"] {
				logrus.WithField("queryContent", queryContent).Info("Hook")
				mywindown.MyWindown.Resize(fyne.NewSize(300, 0))
				if curContent == queryContent {
					mywindown.Show()
					continue
				}

				SetCurContent(queryContent)
				var transalteTool = "youdao"
				result := translate.GetTransalteApp(transalteTool).PostQuery(queryContent)
				switch transalteTool {
				case "youdao":
					if len(result) > 0 {
						transalteResultLab.SetText(result[0])
					}
					if len(result) > 1 {
						transalteExplainsLab.SetText(result[1])
					}
				case "caiyun":
					transalteResult := strings.Join(result, ",")
					transalteResultLab.SetText(transalteResult)
				}

				queryContentLab.SetText(queryContent)

			}
			mywindown.Show()
			continue

		case hook.MouseDown:
			if ev.Button == hook.MouseMap["left"] && preKind == hook.MouseDrag {
				// handleData()
			}
		}
		preKind = ev.Kind
		logrus.WithField("hook: ", ev).Info()
	}
}

func handleData() {
	logrus.Info("handleData start")
	// 读取原来的内容
	oldContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData")
		return
	}
	// 模拟按下 Ctrl 键
	robotgo.KeyTap("c", "ctrl")
	tmpContent, err := robotgo.ReadAll()
	if err != nil {
		logrus.WithError(err).Error("handleData")
		return
	}
	SetQueryContent(tmpContent)
	// 将原来的数据写回去，防止污染剪贴板
	if oldContent != queryContent {
		if err := robotgo.WriteAll(oldContent); err != nil {
			logrus.WithError(err).Error("handleData")
		}
	}
	logrus.Info("handleData end")
	time.Sleep(time.Millisecond * 600)
}
