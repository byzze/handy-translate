package windows

import (
	"log/slog"
	"runtime"
	"syscall"

	"github.com/lxn/win"
)

// GetCursorPos 获取鼠标位置
func GetCursorPos() *win.POINT {
	lpPoint := &win.POINT{}
	win.GetCursorPos(lpPoint)
	return lpPoint
}

func pressKey(keyCode uint16, flag uint) {
	keybdEventProc.Call(uintptr(keyCode), 0, uintptr(flag), 0)
}

func PressCtrlC() {
	// 按下 Ctrl
	pressKey(VK_CONTROL, 0)
	// 按下 C
	pressKey(VK_C, 0)
	// 释放 C
	pressKey(VK_C, KEYEVENTF_KEYUP)
	// 释放 Ctrl
	pressKey(VK_CONTROL, KEYEVENTF_KEYUP)
}

// FindWindow 查找窗口
func FindWindow(windowName string) *Window {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	lpWindowName, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		slog.Error("FindWindow: UTF16 转换失败",
			slog.String("windowName", windowName),
			slog.Any("err", err))
		return nil
	}

	// 查找窗口
	hwnd := win.FindWindow(nil, lpWindowName)
	if hwnd == 0 {
		slog.Warn("FindWindow: 未找到窗口",
			slog.String("windowName", windowName))
		return nil
	}

	w := &Window{
		Name: windowName,
		HWND: hwnd,
	}

	slog.Debug("FindWindow: 成功找到窗口",
		slog.String("windowName", windowName),
		slog.Any("hwnd", hwnd))

	return w
}
