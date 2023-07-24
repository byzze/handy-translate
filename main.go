package main

import (
	"golang/mytheme"
	"golang/register"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 设置字体

	// 鼠标选中内容
	// 模拟按键拷贝数据, 操作完成之后还原剪贴板内容
	// 正则解析内容，调用翻译接口
	// fyne显示翻译内容和原始文本
	// 失去焦点，隐藏窗口

	// 设置自定主题，解决中文乱码
	t := &mytheme.MyTheme{}
	t.SetFonts("./Consolas-with-Yahei Bold Nerd Font.ttf", "")

	a := app.New()

	// 应用自定义主题
	a.Settings().SetTheme(t)

	myWindow := a.NewWindow("翻译工具")
	myWindow.Resize(fyne.NewSize(300, 0))

	myWindow.SetContent(widget.NewLabel("程序启动成功"))

	myWindow.Show()
	go register.Hook(myWindow)

	a.Run()
}
