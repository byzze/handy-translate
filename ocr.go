package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

func ExecOCR(path, image string) string {
	// 指定要执行的可执行文件路径
	executablePath := path

	// 创建一个结构体实例，指定命令和参数
	program := ExternalProgram{
		Command: executablePath,
		Args:    []string{"--image=" + image},
	}

	// 使用结构体中的信息执行外部程序
	output, err := runExternalProgram(program)
	if err != nil {
		logrus.Error("执行外部程序时发生错误：", err)
		return ""
	}
	logrus.Info(string(output))
	// 查找第一个左大括号的位置
	startIndex := strings.Index(string(output), "{")
	if startIndex == -1 {
		logrus.Error("无法找到 JSON 数据的起始位置")
		return ""
	}

	// 从左大括号开始提取 JSON 数据
	jsonStr := output[startIndex:]

	// 解析 JSON 数据
	var result OCRResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		logrus.Error("解析输出结果时发生错误：", err)
		return ""
	}

	var text []string
	// 打印解析后的结果
	for _, item := range result.Data {
		text = append(text, item.Text)
	}

	return strings.Join(text, "\n")
}

type OCRResult struct {
	Code int `json:"code"`
	Data []struct {
		Box   [4][2]int `json:"box"`
		Score float64   `json:"score"`
		Text  string    `json:"text"`
	} `json:"data"`
}

// 定义一个结构体来表示外部程序的命令和参数
type ExternalProgram struct {
	Command string
	Args    []string
}

func runExternalProgram(program ExternalProgram) ([]byte, error) {
	// 使用结构体中的信息创建命令对象
	cmd := exec.Command(program.Command, program.Args...)

	// 设置标准输入、输出和错误输出，如果需要的话
	// 创建一个字节缓冲区来捕获输出
	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	// 启动外部进程
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	// 等待外部进程完成
	err = cmd.Wait()
	return outputBuffer.Bytes(), nil
}

func saveBase64Image(base64String, filename string) error {
	// 将Base64编码的字符串解码为字芴切片
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}

	// 创建一个文件用于保存图片
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入数据到文件
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// 保存Base64字符串到文件（可选）
func saveBase64ToFile(filename, base64Image string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(base64Image)
	if err != nil {
		panic(err)
	}
}
