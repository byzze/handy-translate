package main

import (
	"bytes"
	"context"
	"fmt"
	"handy-translate/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"golang.org/x/sys/windows/registry"
)

func TestConfig(t *testing.T) {
	config.Init(context.TODO())
	config.Save()
}

// func TestOCR(t *testing.T) {
// 	client := gosseract.NewClient()
// 	defer client.Close()
// 	client.SetImage("test.png")
// 	text, _ := client.Text()
// 	fmt.Println(text)
// }

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
