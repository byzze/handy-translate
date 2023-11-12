package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
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
		fmt.Println("执行外部程序时发生错误：", err)
		return ""
	}
	fmt.Println(string(output))
	// 查找第一个左大括号的位置
	startIndex := strings.Index(string(output), "{")
	if startIndex == -1 {
		fmt.Println("无法找到 JSON 数据的起始位置")
		return ""
	}

	// 从左大括号开始提取 JSON 数据
	jsonStr := output[startIndex:]

	// 解析 JSON 数据
	var result OCRResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		fmt.Println("解析输出结果时发生错误：", err)
		return ""
	}

	var text string
	// 打印解析后的结果
	fmt.Printf("Code: %d\n", result.Code)
	for i, item := range result.Data {
		fmt.Printf("Item %d:\n", i+1)
		fmt.Printf("  Box: %v\n", item.Box)
		fmt.Printf("  Score: %f\n", item.Score)
		fmt.Printf("  Text: %s\n", item.Text)
		text += item.Text + "\n"
	}
	return text
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

	// 启动外部进程
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	// 等待外部进程完成
	err = cmd.Wait()
	return outputBuffer.Bytes(), nil
}
