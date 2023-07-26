package xcgui

import (
	"fmt"
	"lyzee-translate/register"
	"lyzee-translate/translate"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var w *window.Window
var st *widget.ShapeText

func Init() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w = window.New(0, 0, 430, 300, "翻译工具", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// w.EnableLayout(true)
	// w.ShowLayoutFrame(true)
	// 设置窗口边框大小
	w.SetBorderSize(0, 30, 0, 0)
	// 设置窗口图标
	a.SetWindowIcon(imagex.NewBySvgStringW(svgIcon).Handle)
	// 设置窗口透明类型
	// w.SetTransparentType(xcc.Window_Transparent_False)
	// 设置窗口阴影
	w.SetShadowInfo(8, 255, 10, false, 0)
	// h1 := xc.XLayoutFrame_Create(0, 0, 100, 100, w.Handle)
	// xc.XWidget_LayoutItem_SetWidth(h1, xcc.Layout_Size_Fill, 0)
	// xc.XWidget_LayoutItem_SetHeight(h1, xcc.Layout_Size_Auto, 0)
	// xc.XLayout_Create(0, 0, 100, 100, w.Handle)
	// xc.XWidget_LayoutItem_SetWidth(h1, xcc.Layout_Size_Fill, 0)
	// xc.XWidget_LayoutItem_SetHeight(h1, xcc.Layout_Size_Weight, 1)
	// 创建按钮
	/* btn := widget.NewButton(165, 135, 100, 30, "Button", w.Handle)
	// 注册按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		a.MessageBox("提示", btn.GetText(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return 0
	}) */

	// widget.NewButton(10, 10, 100, 30, "阿打发打发地方1", h1)
	// widget.NewButton(10, 10, 100, 30, "阿打发打发地方2", h1)
	// widget.NewButton(10, 10, 100, 30, "阿打发打发地方3", h1)
	// widget.NewButton(10, 10, 100, 30, "阿打发打发地方4", h1)
	// widget.NewButton(10, 10, 100, 30, "阿打发打发地方5", h1)
	// widget.NewButton(10, 10, 100, 30, "阿打发打发地方6", h1)

	st = widget.NewShapeText(15, 35, 330, 30, ``, w.Handle)
	st.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Top)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)
	// a.MessageBox("提示", "翻译工具启动成功", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
	// 设置字体大小
	st.SetFont(font.New(12).Handle)

	go handleHookData()
	w.AdjustLayout()
	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

var svgIcon = `<svg t="1669088647057" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5490" width="22" height="22"><path d="M517.12 512.8704m-432.3328 0a432.3328 432.3328 0 1 0 864.6656 0 432.3328 432.3328 0 1 0-864.6656 0Z" fill="#51C5FF" p-id="5491"></path><path d="M292.1472 418.7136c-85.0432 0-160.4096 41.3696-207.104 105.0624 4.5568 182.7328 122.368 337.3056 285.952 396.032 103.2192-33.28 177.92-130.048 177.92-244.3776 0-141.7216-114.944-256.7168-256.768-256.7168z" fill="#7BE0FF" p-id="5492"></path><path d="M800.2048 571.6992l-101.888-58.8288 101.888-58.8288c16.896-9.728 22.6816-31.3344 12.9536-48.2304l-55.296-95.744c-9.728-16.896-31.3344-22.6816-48.2304-12.9536l-101.888 58.8288V238.336c0-19.5072-15.8208-35.328-35.328-35.328H461.824c-19.5072 0-35.328 15.8208-35.328 35.328v117.6064L324.608 297.1136c-16.896-9.728-38.5024-3.9424-48.2304 12.9536l-55.296 95.744c-9.728 16.896-3.9424 38.5024 12.9536 48.2304l101.888 58.8288-101.888 58.8288c-16.896 9.728-22.6816 31.3344-12.9536 48.2304l55.296 95.744c9.728 16.896 31.3344 22.6816 48.2304 12.9536l101.888-58.8288v117.6064c0 19.5072 15.8208 35.328 35.328 35.328h110.592c19.5072 0 35.328-15.8208 35.328-35.328v-117.6064l101.888 58.8288c16.896 9.728 38.5024 3.9424 48.2304-12.9536l55.296-95.744c9.728-16.896 3.9424-38.5024-12.9536-48.2304z" fill="#CAF8FF" p-id="5493"></path><path d="M517.12 512.8704m-234.24 0a234.24 234.24 0 1 0 468.48 0 234.24 234.24 0 1 0-468.48 0Z" fill="#FFFFFF" p-id="5494"></path><path d="M517.12 512.8704m-103.5776 0a103.5776 103.5776 0 1 0 207.1552 0 103.5776 103.5776 0 1 0-207.1552 0Z" fill="#51C5FF" p-id="5495"></path></svg>`

func handleHookData() {
	for {
		select {
		case ev := <-register.HookChan:
			logrus.Info("handleHookData: ", ev)
			curContent := register.GetCurContent()
			queryContent := register.GetQueryContent()
			if curContent == queryContent {
				w.Show(true)
				continue
			}
			register.SetQueryContent(queryContent)
			var text = []string{queryContent}
			var transalteTool = "youdao"
			result := translate.GetTransalteApp(transalteTool).PostQuery(queryContent)
			fmt.Println(result)
			switch transalteTool {
			case "youdao":
				text = append(text, result...)
			case "caiyun":
				text = append(text, result...)
			}
			st.SetText(strings.Join(text, "\n\n"))
			w.Show(true)
		}
	}
}
