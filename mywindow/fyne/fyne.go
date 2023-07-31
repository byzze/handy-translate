package fyne

import (
	"lyzee-translate/register"
	"lyzee-translate/theme"
	"lyzee-translate/translate"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

var myWindown fyne.Window

var myLk sync.Mutex

var queryContentLab, transalteResultLab, transalteExplainsLab *widget.Label

func Run() {
	a := app.New()
	// 设置自定主题，解决中文乱码
	// 应用自定义主题
	a.Settings().SetTheme(&theme.MyTheme{})

	nw := a.NewWindow("翻译工具")
	nw.Resize(fyne.NewSize(400, 0))
	nw.SetMainMenu(nil)
	nw.SetMaster()
	nw.CenterOnScreen()

	queryContentLab = widget.NewLabel("程序启动成功")
	queryContentLab.Wrapping = fyne.TextWrapWord

	transalteResultLab = widget.NewLabel("")
	transalteResultLab.Wrapping = fyne.TextWrapWord

	transalteExplainsLab = widget.NewLabel("")
	transalteExplainsLab.Wrapping = fyne.TextWrapWord

	nw.SetContent(
		container.NewVBox(
			queryContentLab,
			transalteResultLab,
			transalteExplainsLab,
		))

	go handleHookData()

	myWindown = nw
	myWindown.ShowAndRun()
}

// Show 处理数据竞态
func Show() {
	myWindown.Show()
}

// Show 处理数据竞态
func Hide() {
	myWindown.Hide()
}

func handleHookData() {
	for {
		select {
		case <-register.HookCenterChan:
			logrus.Info("handleHookData")
			curContent := register.GetCurContent()
			queryContent := register.GetQueryContent()
			if curContent == queryContent {
				myWindown.Show()
				continue
			}

			register.SetQueryContent(queryContent)
			var transalteTool = "youdao"
			result := translate.GetTransalteApp(transalteTool).PostQuery(queryContent)
			logrus.WithField("result", result).Info("Transalte")
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
			myWindown.Show()
		}
	}
}
