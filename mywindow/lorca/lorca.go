package lorca

import (
	"bytes"
	"handy-translate/config"
	"handy-translate/register"
	"handy-translate/translate"
	"handy-translate/translate/caiyun"
	"handy-translate/translate/youdao"
	"html/template"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"github.com/sirupsen/logrus"
	"github.com/zserge/lorca"
)

type transalte struct {
	Title        string
	QueryContent string
	Explain      string
	ExplainEx    string
}

var t *transalte
var tmpl *template.Template
var ui lorca.UI

// Run lorca runing
func Run() {
	var err error
	var width, height = 300, 400
	t = &transalte{
		Title:        config.Data.Appname,
		QueryContent: "程序启动成功",
	}
	var b bytes.Buffer
	tmpl, err = template.ParseFiles("mywindow/lorca/index.html")
	tmpl.Execute(&b, t)
	content := b.String()
	ui, err = lorca.New("data:text/html,"+url.PathEscape(content), "", width, height, "--remote-allow-origins=*")
	if err != nil {
		logrus.Panic(err)
	}

	go processData()

	defer ui.Close()
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	logrus.Info("exiting...")
}

func processData() {
	for {
		select {
		case <-register.HookCenterChan:
			logrus.Info("processData")
			curContent := register.GetCurText()
			text := register.GetQueryText()
			if curContent == text {
				show()
				continue
			}

			register.SetCurText(text)
			t.QueryContent = text

			var transalteTool = config.Data.TranslateWay
			way := translate.GetTransalteWay(transalteTool)
			result := way.PostQuery(text)
			logrus.WithField("result", result).Info("Transalte")

			switch way.(type) {
			case *youdao.Youdao:
				t.Explain = result[0]
				t.ExplainEx = result[1]
			case *caiyun.Caiyun:
				t.QueryContent = text
				t.Explain = strings.Join(result, ",")
			}

			var b bytes.Buffer
			tmpl.Execute(&b, t)
			content := b.String()
			loadableContents := "data:text/html," + url.PathEscape(content)
			ui.Load(loadableContents)

			show()
		}
	}
}

func show() {
	lpWindowName, err := syscall.UTF16PtrFromString(config.Data.Appname)
	if err != nil {
		logrus.WithError(err).Error("UTF16PtrFromString")
	}

	// find window
	hwnd := win.FindWindow(nil, lpWindowName)
	if hwnd == 0 {
		logrus.Panic("启动失败")
		return
	}

	var rect win.RECT
	win.GetWindowRect(hwnd, &rect)
	width := rect.Right - rect.Left
	height := rect.Bottom - rect.Top

	x, y := robotgo.GetMousePos()
	win.SetWindowPos(hwnd, 0, int32(x), int32(y), width, height, win.SWP_SHOWWINDOW)
	win.SetForegroundWindow(hwnd)
}
