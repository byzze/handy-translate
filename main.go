package main

import (
	"fmt"
	"image/color"
	"io"
	"lyzee-translate/mywindown/xcgui"
	"lyzee-translate/register"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// 判断颜色是否属于文字的颜色范围
func isTextColor(c color.RGBA) bool {
	// 在这里根据你关注的文字颜色范围进行判断
	// 示例：假设文字颜色范围是红色到蓝色之间
	if c.R >= 0 && c.R <= 255 && c.G == 0 && c.B == 0 {
		return true
	}
	return false
}

// 将十六进制颜色代码转换为 color.RGBA
func hexToRGBA(hexColor string) color.RGBA {
	// 将十六进制颜色代码转换为对应的整数值
	var intValue uint32
	fmt.Sscanf(hexColor, "%x", &intValue)

	// 提取红色、绿色和蓝色通道的值
	red := uint8((intValue >> 16) & 0xFF)
	green := uint8((intValue >> 8) & 0xFF)
	blue := uint8(intValue & 0xFF)

	// 创建 color.RGBA 实例
	rgbaColor := color.RGBA{R: red, G: green, B: blue, A: 255}

	return rgbaColor
}

func main() {
	/* fmt.Println("start")
	for {
		event := robotgo.AddEvent("mleft")
		if event {
			fmt.Println("检测到鼠标点击")
			x, y := robotgo.GetMousePos()
			colorVal := robotgo.GetPixelColor(x, y)
			rgbaColor := hexToRGBA(colorVal)
			fmt.Println(rgbaColor)
			if isTextColor(rgbaColor) {
				fmt.Println("点击的是文字")
			} else {
				fmt.Println("点击的不是文字")
			}
		}
	} */
	// 设置字体

	// 鼠标选中内容
	// 模拟按键拷贝数据, 操作完成之后还原剪贴板内容
	// 正则解析内容，调用翻译接口
	// fyne显示翻译内容和原始文本
	// 失去焦点，隐藏窗口

	var tnow = time.Now()
	tnowStr := tnow.Format("20060102")
	// 校验文件是否存在
	filename := tnowStr + ".log"
	_, err := os.Stat(filename)
	var file *os.File
	if os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err = os.Create(filename)
		if err != nil {
			fmt.Println("无法创建文件:", err)
			return
		}
		fmt.Println("文件创建成功。")
	} else {
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("无法创建文件:", err)
			return
		}
	}

	defer file.Close()

	file.Seek(0, 0) //TODO 每次运行清空日志
	mw := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(mw)

	go register.Hook()

	// gio.Init()
	// gkt3.Init()
	xcgui.Init()
	// fynewindown.Init()
}
