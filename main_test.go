package main

import (
	"bytes"
	"context"
	"fmt"
	"handy-translate/config"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/kbinani/screenshot"
	"golang.org/x/sys/windows/registry"
)

var bs = `data:image/png;base64,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`

func TestConfig(t *testing.T) {
	config.Init(context.TODO())
	config.Save()
}

//	func TestOCR(t *testing.T) {
//		client := gosseract.NewClient()
//		defer client.Close()
//		client.SetImage("test.png")
//		text, _ := client.Text()
//		fmt.Println(text)
//	}
//
// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func TestOCR2(t *testing.T) {
	// 获取屏幕截图
	// i := 0
	// bounds := screenshot.GetDisplayBounds(i)

	// img, err := screenshot.CaptureRect(bounds)
	// if err != nil {
	// 	panic(err)
	// }
	// fileName := "./frontend/screenshot/screenshot-1699605146407919600.png"
	// file, _ := os.Create(fileName)
	// defer file.Close()
	// png.Encode(file, img)

	// fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	img, err := screenshot.Capture(423, 198, 531, 92)
	if err != nil {
		panic(err)
	}
	save(img, "all.png")

	// {
	// 	"left": 423,
	// 	"top": 198,
	// 	"width": 531,
	// 	"height": 92
	// }
}

func TestAutoStarup(t *testing.T) {
	// 注册表路径
	keyPath := `Software\Microsoft\Windows\CurrentVersion\Run`

	// 要设置的键名和值（你的程序的路径）
	valueName := "MyGolangApp"
	valueData := `C:\Users\loyd\Desktop\byzze\handy-translate-install\handy-translate\handy-translate.exe`

	// 打开或创建注册表项
	k, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.WRITE)
	if err != nil {
		fmt.Println("Error opening or creating registry key:", err)
		os.Exit(1)
	}
	defer k.Close()

	// 设置注册表项的值
	err = k.SetStringValue(valueName, valueData)
	if err != nil {
		fmt.Println("Error setting registry key value:", err)
		os.Exit(1)
	}

	fmt.Println("Registry key created and set successfully.")
}

func TestNotAutoStarup(t *testing.T) {
	// 打开注册表项
	keyPath := `Software\Microsoft\Windows\CurrentVersion\Run`

	// 要设置的键名和值（你的程序的路径）
	valueName := "MyGolangApp"

	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		return
	}
	defer key.Close()

	// 删除注册表项中的相应值
	if err := key.DeleteValue(valueName); err != nil {
		fmt.Println("Error deleting registry value:", err)
		return
	}

	fmt.Println("Startup entry removed for YourAppName")
}

func TestPingRoute(t *testing.T) {
	url := "https://dict.youdao.com/suggest?num=5&ver=3.0&doctype=json&cache=false&le=en&q=hello" // 替换为你要请求的 URL

	// 发起 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP 请求出错:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应内容出错:", err)
		return
	}

	fmt.Println("响应内容:", string(body))
}

func TestLangdetect(t *testing.T) {
	// 目标 URL
	uri := "https://fanyi.baidu.com/langdetect"

	// 准备请求体数据
	data := url.Values{}
	data.Set("query", "asd")
	payload := bytes.NewBufferString(data.Encode())

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", uri, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 输出响应数据
	fmt.Println(string(body))

}
