package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应内容出错:", err)
		return
	}

	fmt.Println("响应内容:", string(body))

}
