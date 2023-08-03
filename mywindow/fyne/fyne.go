package fyne

import (
	"handy-translate/mywindow/fyne/theme"
	"handy-translate/register"
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

	a.Settings().SetTheme(&theme.MyTheme{})

	nw := a.NewWindow("handy-translate")
	nw.Resize(fyne.NewSize(400, 0))
	nw.SetMainMenu(nil)
	nw.SetMaster()
	nw.CenterOnScreen()

	queryContentLab = widget.NewLabel("start success")
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

	go processData()

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

func processData() {
	for {
		select {
		case <-register.HookCenterChan:
			logrus.Info("processData")
			curContent := register.GetCurText()
			queryContent := register.GetQueryText()
			if curContent == queryContent {
				myWindown.Show()
				continue
			}

			register.SetQueryContent(queryContent)

			queryContentLab.SetText(queryContent)
			myWindown.Show()
		}
	}
}
