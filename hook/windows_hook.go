package hook

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
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
	user32              = windows.NewLazySystemDLL("user32.dll")
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

func LowLevelMouseProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	r1, _, _ := callNextHookEx.Call(hHook, uintptr(nCode), wParam, lParam)
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
			ch <- struct{}{}
			if len(ch) >= 2 {
				for {
					select {
					case <-ch:
						// 处理通道的值
					default:
						// 通道为空时执行的代码
						return r1
					}
				}
			}
			fmt.Println("中键按下")
			copyText()
			/* textToCopy := "这是要复制到剪贴板的文本"
			err := setClipboardText(textToCopy)
			if err != nil {
				fmt.Println("复制到剪贴板时发生错误:", err)
			} else {
				fmt.Println("文本已复制到剪贴板")
			}

			text, err := getClipboardText()
			if err != nil {
				fmt.Println("从剪贴板粘贴文本时发生错误:", err)
			} else {
				fmt.Println("从剪贴板粘贴的文本:", text)
			} */
		case WM_MBUTTONUP:
			// fmt.Println("中键释放")
		}
	}
	return r1
}

func WindowsHook() {
	hInstance := windows.Handle(0) // 在Go中无需使用 GetModuleHandle

	hHook, _, _ = setWindowsHookEx.Call(uintptr(WH_MOUSE_LL), syscall.NewCallback(LowLevelMouseProc), uintptr(hInstance), 0)

	defer unhookWindowsHookEx.Call(hHook)

	text, _ := syscall.UTF16PtrFromString("这是一个消息框示例")
	for {
		// 监听消息
		getMessage.Call(uintptr(unsafe.Pointer(&text)), 0, 0, 0)
	}
}

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	openClipboard    = user32.NewProc("OpenClipboard")
	closeClipboard   = user32.NewProc("CloseClipboard")
	emptyClipboard   = user32.NewProc("EmptyClipboard")
	getClipboardData = user32.NewProc("GetClipboardData")
	setClipboardData = user32.NewProc("SetClipboardData")
	globalAlloc      = kernel32.NewProc("GlobalAlloc")
	globalFree       = kernel32.NewProc("GlobalFree")
	globalLock       = kernel32.NewProc("GlobalLock")
	globalUnlock     = kernel32.NewProc("GlobalUnlock")
	copyMemory       = kernel32.NewProc("RtlCopyMemory")
)

const (
	cfUnicodeText = 13
	gmMoveable    = 0x0002
	ghnd          = 0x0042
)

func openClipboardFunc(hWnd uintptr) bool {
	r, _, _ := openClipboard.Call(hWnd)
	return r != 0
}

func closeClipboardFunc() bool {
	r, _, _ := closeClipboard.Call()
	return r != 0
}

func emptyClipboardFunc() bool {
	r, _, _ := emptyClipboard.Call()
	return r != 0
}

func setClipboardText(text string) error {
	if !openClipboardFunc(0) {
		return fmt.Errorf("无法打开剪贴板")
	}
	defer closeClipboardFunc()

	if !emptyClipboardFunc() {
		return fmt.Errorf("无法清空剪贴板")
	}

	utf16Text, _ := syscall.UTF16FromString(text)

	hGlobal, _, _ := globalAlloc.Call(ghnd, uintptr((len(utf16Text)+1)*2))
	if hGlobal == 0 {
		return fmt.Errorf("无法分配全局内存")
	}
	defer globalFree.Call(hGlobal)

	lpGlobalLock, _, _ := globalLock.Call(hGlobal)
	if lpGlobalLock == 0 {
		return fmt.Errorf("无法锁定全局内存")
	}

	copyMemory.Call(lpGlobalLock, uintptr(unsafe.Pointer(&utf16Text[0])), uintptr((len(utf16Text)+1)*2))
	globalUnlock.Call(hGlobal)

	r, _, _ := setClipboardData.Call(cfUnicodeText, hGlobal)
	if r == 0 {
		return fmt.Errorf("无法设置剪贴板数据")
	}

	return nil
}

func getClipboardText() (string, error) {
	if !openClipboardFunc(0) {
		return "", fmt.Errorf("无法打开剪贴板")
	}
	defer closeClipboardFunc()

	hData, _, _ := getClipboardData.Call(cfUnicodeText)
	if hData == 0 {
		return "", fmt.Errorf("无法获取剪贴板数据")
	}

	lpGlobalLock, _, _ := globalLock.Call(hData)
	if lpGlobalLock == 0 {
		return "", fmt.Errorf("无法锁定全局内存")
	}
	defer globalUnlock.Call(hData)

	text := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(lpGlobalLock))[:])
	return text, nil
}

var (
	user321     = syscall.MustLoadDLL("user32.dll")
	keybd_event = user321.MustFindProc("keybd_event")
	// getForegroundWindow = user321.MustFindProc("GetForegroundWindow")
	// sendMessage         = user321.MustFindProc("SendMessageW")
)

const (
	// 定义键盘事件常量
	KEYEVENTF_KEYDOWN = 0x0000
	KEYEVENTF_KEYUP   = 0x0002

	// 定义VK常量
	VK_CONTROL = 0x11
	VK_C       = 0x43
)

func keybdEvent(keycode, flags uintptr) {
	keybd_event.Call(keycode, uintptr(0), flags, 0)
}

func copyText() error {
	// 获取当前活动窗口的句柄
	// hwnd, _, _ := getForegroundWindow.Call()
	// if hwnd == 0 {
	// 	return fmt.Errorf("无法获取当前活动窗口句柄")
	// }

	// 模拟按下Ctrl键
	keybdEvent(VK_CONTROL, KEYEVENTF_KEYDOWN)
	time.Sleep(100 * time.Millisecond) // 等待一会
	defer keybdEvent(VK_CONTROL, KEYEVENTF_KEYUP)

	// 模拟按下C键
	keybdEvent(VK_C, KEYEVENTF_KEYDOWN)
	time.Sleep(100 * time.Millisecond) // 等待一会
	defer keybdEvent(VK_C, KEYEVENTF_KEYUP)

	// 使用SendMessage来发送WM_COPY消息，将复制命令发送到当前活动窗口
	// result, _, err := sendMessage.Call(hwnd, 0x0301, 0, 0)
	// fmt.Println(err)
	// if result == 0 {
	// 	return fmt.Errorf("无法发送复制命令")
	// }

	return nil
}
