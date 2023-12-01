package windows

import (
	"log/slog"
	"runtime"
	"syscall"

	"github.com/lxn/win"
)

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
