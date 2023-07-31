package lorca

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zserge/lorca"
)

var UI lorca.UI

// 导入Windows API函数
var (
	user32              = syscall.NewLazyDLL("user32.dll")
	showWindowProc      = user32.NewProc("ShowWindow")
	getForegroundWin    = user32.NewProc("GetForegroundWindow")
	setForegroundWindow = user32.NewProc("SetForegroundWindow")
)

const (
	swHide  = 0
	SW_HIDE = 0
	swShow  = 5
)

func showWindow(hwnd syscall.Handle, showCmd int) {
	showWindowProc.Call(uintptr(hwnd), uintptr(showCmd))
}

// 隐藏窗口
func HideWindow(hwnd syscall.Handle) {
	// 使用ShowWindow函数隐藏窗口
	_, _, err := showWindowProc.Call(uintptr(hwnd), uintptr(0)) // 0表示隐藏窗口
	if err != nil && err.Error() != "The operation completed successfully." {
		log.Fatal(err)
	}
}

// 显示窗口
func ShowWindow(hwnd syscall.Handle) {
	// 使用ShowWindow函数显示窗口
	_, _, err := showWindowProc.Call(uintptr(hwnd), uintptr(1)) // 1表示显示窗口
	if err != nil && err.Error() != "The operation completed successfully." {
		log.Fatal(err)
	}
}

func Run() {

	var err error
	UI, err = lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 300, 300, "--remote-allow-origins=*")

	if err != nil {
		log.Fatal(err)
	}
	// 获取当前窗口句柄
	// 获取前台窗口句柄
	foregroundWin, _, _ := getForegroundWin.Call()
	// 获取窗口句柄
	// hwnd := GetHwnd()

	// 显示窗口
	// ShowWindow(hwnd)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			// 隐藏窗口
			showWindowProc.Call(foregroundWin, SW_HIDE)
			time.Sleep(time.Second * 3)
			showWindowProc.Call(foregroundWin, 1)
			setForegroundWindow.Call(foregroundWin)
			fmt.Println("窗口已隐藏")
			// l, t := robotgo.GetMousePos()
			// bd := lorca.Bounds{
			// 	Left:        l,
			// 	Top:         t,
			// 	WindowState: lorca.WindowStateNormal,
			// }
			// UI.SetBounds(bd)
			// logrus.Info(bd)
		}
	}()

	defer UI.Close()
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-UI.Done():
	}

	log.Println("exiting...")
}
