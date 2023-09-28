package hook

import (
	"syscall"
	"unsafe"

	"github.com/go-vgo/robotgo"
)

const (
	WH_MOUSE_LL    = 14
	WM_MOUSEMOVE   = 0x0200
	WM_LBUTTONDOWN = 0x0201
	WM_LBUTTONUP   = 0x0202
	WM_RBUTTONDOWN = 0x0204
	WM_RBUTTONUP   = 0x0205
	WM_MBUTTONDOWN = 0x0207
	WM_MBUTTONUP   = 0x0208
	MK_LBUTTON     = 0x0001
	MK_RBUTTON     = 0x0002
	MK_MBUTTON     = 0x0010
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	callNextHookEx      = user32.NewProc("CallNextHookEx")
	unhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	getMessage          = user32.NewProc("GetMessageW")
)

type MSLLHOOKSTRUCT struct {
	Pt          POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

type POINT struct {
	X, Y int32
}

var (
	hHook uintptr
)
var ch = make(chan struct{}, 2)

// LowLevelMouseProc 代用windows api 才能做到选中文字，鼠标事件触发前执行模拟ctrl + c 操作
func LowLevelMouseProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	r1, _, _ := callNextHookEx.Call(uintptr(nCode), wParam, lParam)
	if nCode >= 0 {
		// msllHookStruct := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		switch wParam {
		case WM_MOUSEMOVE:
			// fmt.Println("鼠标移动：", msllHookStruct.Pt.X, msllHookStruct.Pt.Y)
		case WM_LBUTTONDOWN:
			// fmt.Println("左键按下")
		case WM_LBUTTONUP:
			// fmt.Println("左键释放")
		case WM_RBUTTONDOWN:
			// fmt.Println("右键按下")
		case WM_RBUTTONUP:
			// fmt.Println("右键释放")
		case WM_MBUTTONDOWN:
			robotgo.KeyTap("c", "ctrl")
			// copyText()
		case WM_MBUTTONUP:
		}
	}
	return r1
}

func WindowsHook() {

	hHook, _, _ = setWindowsHookEx.Call(uintptr(WH_MOUSE_LL), syscall.NewCallback(LowLevelMouseProc), 0, 0)

	defer unhookWindowsHookEx.Call(hHook)

	str, _ := syscall.UTF16PtrFromString("")

	// 阻塞消息
	for {
		// 监听消息
		getMessage.Call(uintptr(unsafe.Pointer(&str)), 0, 0, 0)
	}
}
