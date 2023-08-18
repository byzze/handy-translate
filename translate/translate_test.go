package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/OwO-Network/gdeeplx"
	"github.com/sirupsen/logrus"
)

func TestGetTransalteWay(t *testing.T) {
	result, err := gdeeplx.Translate("hello", "EN", "ZH", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)
}

func TestLingva(t *testing.T) {
	url := "https://lingva.1link.fun/_next/data/dK-OKex2AGbb_H1CfdJwC/auto/zh/" + url.QueryEscape("Alternative front-end for Google Translate, serving as a Free and Open Source translator with over a hundred languages available s") + ".json"

	// 替换为你要请求的 URL

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
	var tr DefaultTransalte

	err = json.Unmarshal(body, &tr)
	if err != nil {
		logrus.WithError(err).Error("Unmarshal")
	}

	// fmt.Println("响应内容:", string(body))
	fmt.Println("响应内容:", tr.PageProps.Translation)
}
