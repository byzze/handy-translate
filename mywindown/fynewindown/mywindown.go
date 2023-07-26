package fynewindown

import (
	"fmt"
	"lyzee-translate/register"
	"lyzee-translate/theme"
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

func Init() {
	a := app.New()
	// 设置自定主题，解决中文乱码
	// 应用自定义主题
	a.Settings().SetTheme(&theme.MyTheme{})

	nw := a.NewWindow("翻译工具")
	nw.Resize(fyne.NewSize(300, 0))
	nw.SetMainMenu(nil)
	nw.SetMaster()

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
	// MyWindown.Resize(fyne.NewSize(300, 0))
	fmt.Println("a ...any")
	// if myLk.TryLock() {
	fmt.Println("111")
	myWindown.Show()
	// myLk.Unlock()
	fmt.Println("222")
	// }
}

// Show 处理数据竞态
func Hide() {
	// MyWindown.Resize(fyne.NewSize(300, 0))
	fmt.Println("a ...any")
	// if myLk.TryLock() {
	fmt.Println("111")
	myWindown.Hide()
	// myLk.Unlock()
	fmt.Println("222")
}

func handleHookData() {
	for {
		select {
		case ev := <-register.HookChan:
			logrus.Info("handleHookData: ", ev)
			curContent := register.GetCurContent()
			queryContent := register.GetQueryContent()
			if curContent == queryContent {
				myWindown.Show()
				continue
			}

			register.SetQueryContent(queryContent)
			// var transalteTool = "youdao"
			// result := translate.GetTransalteApp(transalteTool).PostQuery(queryContent)
			// switch transalteTool {
			// case "youdao":
			// 	if len(result) > 0 {
			// 		transalteResultLab.SetText(result[0])
			// 	}
			// 	if len(result) > 1 {
			// 		transalteExplainsLab.SetText(result[1])
			// 	}
			// case "caiyun":
			// 	transalteResult := strings.Join(result, ",")
			// 	transalteResultLab.SetText(transalteResult)
			// }

			queryContentLab.SetText(queryContent)
			myWindown.Show()
		}
	}
}
