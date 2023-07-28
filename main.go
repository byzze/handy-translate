package main

import (
	"fmt"
	"io"
	"lyzee-translate/mywindown/fynewindown"
	"lyzee-translate/register"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

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

	// xcgui.Init()
	fynewindown.Init()
}
