package youdao

import (
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

const YaoDaoOnlineWay = "youdao_online"

type YoudaoOnline struct {
	config.Translate
}

func (y *YoudaoOnline) PostQuery(query string) []string {
	url := "https://dict.youdao.com/suggest?num=2&ver=3.0&doctype=json&cache=false&le=en&q=" + url.QueryEscape(query) // 替换为你要请求的 URL

	// 发起 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP 请求出错:", err)
		return nil
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应内容出错:", err)
		return nil
	}

	var tr YoudaoOnlineTransalte

	err = json.Unmarshal(body, &tr)
	if err != nil {
		slog.Error("Unmarshal", err)
		return nil
	}
	if len(tr.Data.Entries) > 0 {
		return []string{tr.Data.Entries[0].Explain}
	}
	return nil
}

type YoudaoOnlineTransalte struct {
	Result Result `json:"result"`
	Data   Data   `json:"data"`
}

type Result struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Entries struct {
	Explain string `json:"explain"`
	Entry   string `json:"entry"`
}

type Data struct {
	Entries  []Entries `json:"entries"`
	Query    string    `json:"query"`
	Language string    `json:"language"`
	Type     string    `json:"type"`
}
