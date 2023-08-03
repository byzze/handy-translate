package windows

import (
	"handy-translate/config"
	"sync"
	"syscall"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"github.com/sirupsen/logrus"
)

var so sync.Once
var hwnd win.HWND

type Windows struct{}

func (w *Windows) Show() {
	hwnd = getHwnd()
	if hwnd == 0 {
		logrus.Error("获取失败")
	}
	var rect win.RECT
	win.GetWindowRect(hwnd, &rect)
	width := rect.Right - rect.Left
	height := rect.Bottom - rect.Top
	xp, yp := robotgo.GetMousePos()
	xs := win.GetSystemMetrics(win.SM_CXSCREEN)
	ys := win.GetSystemMetrics(win.SM_CYSCREEN)

	x := int32(xp)
	y := int32(yp)

	if y+height > ys {
		y = ys - height - 10
	}
	if x+width > xs {
		x = xs - width - 10
	}

	win.SetForegroundWindow(hwnd)
	win.SetWindowPos(hwnd, 0, int32(x), int32(y), width, height, win.SWP_SHOWWINDOW)
	win.ShowWindow(hwnd, win.SW_RESTORE|win.SWP_FRAMECHANGED|win.SWP_NOMOVE|win.SWP_NOSIZE)
}

func getHwnd() win.HWND {
	so.Do(
		func() {
			lpWindowName, err := syscall.UTF16PtrFromString(config.Data.Appname)
			if err != nil {
				logrus.WithError(err).Error("UTF16PtrFromString")
			}

			lpClassName, err := syscall.UTF16PtrFromString("Chrome_WidgetWin_1")
			if err != nil {
				logrus.WithError(err).Error("UTF16PtrFromString")
			}

			hwnd = win.FindWindow(lpClassName, lpWindowName)
		})
	return hwnd
}
