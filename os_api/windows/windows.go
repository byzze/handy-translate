package windows

import (
	"log/slog"
	"runtime"
	"syscall"

	"github.com/lxn/win"
)

// Window Windows os 窗口
type Window struct {
	Name string
	HWND win.HWND
}

// FindWindow 查找窗口
func FindWindow(windowName string) *Window {
	w := new(Window)
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	lpWindowName, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		slog.Error("UTF16PtrFromString", err)
		return nil
	}

	// find window
	hwnd := win.FindWindow(nil, lpWindowName)
	if hwnd == 0 {
		slog.Error("FindWindow Failed")
		return nil
	}
	w.HWND = hwnd
	return w
}

// ShowForWindows windows 窗口下的弹窗， 因为wails的弹窗无法和通过鼠标有效弹出，这里采用windows原生api
func (w Window) ShowForWindows() {
	hwnd := w.HWND
	win.SetForegroundWindow(hwnd)
	win.ShowWindow(hwnd, win.SW_SHOW)
}

// GetCursorPos 获取鼠标位置
func GetCursorPos() *win.POINT {
	lpPoint := &win.POINT{}
	win.GetCursorPos(lpPoint)
	return lpPoint
}
