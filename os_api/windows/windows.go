package windows

import (
	"log/slog"

	"github.com/lxn/win"
)

// Window Windows os 窗口
type Window struct {
	Name string
	HWND win.HWND
}

// ShowForWindows windows 窗口下的弹窗， 因为wails的弹窗无法和通过鼠标有效弹出，这里采用windows原生api
func (w *Window) ShowForWindows() {
	if w == nil {
		slog.Error("ShowForWindows: 窗口对象为 nil")
		return
	}

	if w.HWND == 0 {
		slog.Error("ShowForWindows: 无效的窗口句柄",
			slog.String("windowName", w.Name))
		return
	}

	hwnd := w.HWND

	// 设置为前台窗口
	if !win.SetForegroundWindow(hwnd) {
		slog.Warn("ShowForWindows: SetForegroundWindow 失败",
			slog.String("windowName", w.Name))
	}

	// 显示窗口
	if !win.ShowWindow(hwnd, win.SW_SHOW) {
		slog.Warn("ShowForWindows: ShowWindow 失败",
			slog.String("windowName", w.Name))
	} else {
		slog.Debug("ShowForWindows: 成功显示窗口",
			slog.String("windowName", w.Name))
	}
}
