package windows

import (
	"log/slog"
	"runtime"
	"syscall"

	"github.com/lxn/win"
)

type Window struct {
	Name string
	HWND win.HWND
}

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
	win.ShowWindow(hwnd, win.SW_SHOW)
}

// ShowForWindows windows 窗口下的弹窗， 因为wails的弹窗无法和通过鼠标有效弹出，这里采用windows原生api
func ShowForWindows(windowName string) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	lpWindowName, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		slog.Error("UTF16PtrFromString", err)
	}

	// find window
	hwnd := win.FindWindow(nil, lpWindowName)
	if hwnd == 0 {
		slog.Error("FindWindow Failed")
		return
	}
	win.SetForegroundWindow(hwnd)
	win.ShowWindow(hwnd, win.SW_SHOW)
}

/*
func (w Window) FullScreen() {
	hwnd := w.HWND
	win.SetForegroundWindow(hwnd)
	win.ShowWindow(hwnd, win.SW_SHOW)
}

 var (
	findWindow       = user32.NewProc("FindWindowW")
	setWindowPos     = user32.NewProc("SetWindowPos")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

const (
	SWP_SHOWWINDOW   = 0x0040
	SWP_FRAMECHANGED = 0x0020
)

func setWindowFullScreen(hwnd syscall.Handle) {
	setWindowPos.Call(
		uintptr(hwnd),
		uintptr(0),
		0,
		0,
		uintptr(GetSystemMetrics(0)),
		uintptr(GetSystemMetrics(1)),
		SWP_SHOWWINDOW|SWP_FRAMECHANGED,
	)
}

func GetSystemMetrics(nIndex int) int {
	ret, _, _ := syscall.Syscall(getSystemMetrics.Addr(), 1, uintptr(nIndex), 0, 0)
	return int(ret)
}
*/
