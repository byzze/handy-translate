//go:build windows

package toolbar

import (
	"syscall"
	"unsafe"
)

var (
	user32                = syscall.NewLazyDLL("user32.dll")
	procGetWindowLongPtrW = user32.NewProc("GetWindowLongPtrW")
	procSetWindowLongPtrW = user32.NewProc("SetWindowLongPtrW")
	procSetWindowPos      = user32.NewProc("SetWindowPos")
)

const (
	WS_EX_TOOLWINDOW = 0x00000080
	WS_EX_APPWINDOW  = 0x00040000
	HWND_TOPMOST     = ^uintptr(0)
	SWP_NOMOVE       = 0x0002
	SWP_NOSIZE       = 0x0001
	SWP_NOACTIVATE   = 0x0010
	SWP_FRAMECHANGED = 0x0020
)

var (
	GWL_EXSTYLE = ^uintptr(20 - 1) // -20 in signed, converted to uintptr
)

// setupWindowsToolWindowStyle 设置窗口为工具窗口样式，不在任务栏显示
func setupWindowsToolWindowStyle() {
	// 通过窗口标题查找窗口句柄
	hwnd := findWindowByTitle(WindowName)
	if hwnd != 0 {
		setToolWindowStyle(hwnd)
	}
}

// findWindowByTitle 通过窗口标题查找窗口句柄
func findWindowByTitle(title string) uintptr {
	findWindow := user32.NewProc("FindWindowW")
	ret, _, _ := findWindow.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
	)
	return ret
}

// setToolWindowStyle 设置窗口样式为工具窗口
func setToolWindowStyle(hwnd uintptr) {
	// 获取当前扩展样式
	ret, _, _ := procGetWindowLongPtrW.Call(hwnd, GWL_EXSTYLE)
	exStyle := ret

	// 添加 WS_EX_TOOLWINDOW（不显示在任务栏）
	// 移除 WS_EX_APPWINDOW（不强制显示在任务栏）
	exStyle |= WS_EX_TOOLWINDOW
	exStyle &^= WS_EX_APPWINDOW

	// 设置新的扩展样式
	procSetWindowLongPtrW.Call(hwnd, GWL_EXSTYLE, exStyle)

	// 刷新窗口以应用新样式
	procSetWindowPos.Call(
		hwnd,
		0,
		0, 0, 0, 0,
		SWP_NOMOVE|SWP_NOSIZE|SWP_NOACTIVATE|SWP_FRAMECHANGED,
	)
}
