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
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/registry"
)

var bs = "data:image/png;base64,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"

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

func TestImageBase64(t *testing.T) {
	i := 0
	bounds := screenshot.GetDisplayBounds(i)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	// 将图像编码为Base64字符串
	base64Image := encodeImageToBase64(img)

	// 保存Base64字符串到文件（可选）
	saveBase64ToFile("screenshot_base64.txt", base64Image)

	// // 打印Base64字符串
	println("Base64 Image:")
	println(base64Image)
}

func TestOCR2(t *testing.T) {
	// 获取屏幕截图
	i := 0
	bounds := screenshot.GetDisplayBounds(i)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := "./frontend/screenshot/screenshot-1699605146407919600.png"
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	// fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	// img, err := screenshot.Capture(423, 198, 531, 92)
	// if err != nil {
	// 	panic(err)
	// }
	// save(img, "all.png")

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

func TestParse(t *testing.T) {

	base64String := bs
	base64String = strings.TrimPrefix(base64String, "data:image/png;base64,")
	logrus.Info(base64String)
	// filename := "output.png" // 保存的文件名

	// err := saveBase64Image(base64String, filename)
	// if err != nil {
	// 	log.Fatal("保存图片出错: ", err)
	// }

	log.Println("图片保存成功")
}

func TestOCR(t *testing.T) {
	// 设置可执行文件路径和参数
	executablePath := ".\\ocr\\RapidOCR-json.exe"
	// imagePath := "test.png"
	// arg := fmt.Sprintf("--models=.\\ocr\\models --image=%s", imagePath)

	// 创建一个*Cmd，表示要执行的命令
	cmd := exec.Command(executablePath, "-h")

	// 执行命令并等待完成
	output, err := cmd.Output()

	if err != nil {
		log.Fatalf("命令执行失败：%v", err)
	}

	// 输出命令的结果
	fmt.Println(string(output))
}

func TestExec(t *testing.T) {
	ExecOCR(".\\RapidOCR-json.exe", ".\\output.png")
}
