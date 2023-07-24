package register

import (
	"fmt"
	"golang/youdao/transalte"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

// Hook register hook event
func Hook(myWindow fyne.Window) {
	fmt.Println("--- Please wait hook starting ---")
	var newContent string
	var oldContent string
	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] && e.Clicks == 2 || e.Clicks == 3 {
			fmt.Println(e)
			// 读取原来的内容
			oldContent, err := robotgo.ReadAll()
			if err != nil {
				fmt.Println(err)
				return
			}
			// 模拟按下 Ctrl 键
			robotgo.KeyTap("c", "ctrl")
			newContent, err = robotgo.ReadAll()
			if err != nil {
				fmt.Println(err)
				return
			}
			// 将原来的数据写回去，防止污染剪贴板
			if oldContent != newContent {
				if err := robotgo.WriteAll(oldContent); err != nil {
					fmt.Println(err)
				}
				oldContent = ""
			}
		}
		if e.Button == hook.MouseMap["center"] {
			fmt.Println(newContent)
			myWindow.Resize(fyne.NewSize(300, 0))
			if oldContent == newContent {
				myWindow.Show()
				return
			}
			// 调用翻译接口, 将翻译结果set
			result := transalte.PostQuery(newContent)

			transalteResult := strings.Join(result.Translation, ",")
			transalteExplains := strings.Join(result.Basic.Explains, ",")

			transalteResultLab := widget.NewLabel(transalteResult)
			transalteResultLab.Wrapping = fyne.TextWrapWord

			transalteExplainsLab := widget.NewLabel(transalteExplains)
			transalteExplainsLab.Wrapping = fyne.TextWrapWord

			myWindow.SetContent(
				container.NewVBox(
					widget.NewLabel(newContent),
					transalteResultLab,
					transalteExplainsLab,
				))
			myWindow.Show()
			time.Sleep(time.Millisecond * 300)
		}
	})
	s := hook.Start()
	fmt.Println("--- Please wait hook start success ---")
	<-hook.Process(s)
}
