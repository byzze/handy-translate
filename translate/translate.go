package translate

import (
	"encoding/json"
	"handy-translate/config"
	"handy-translate/translate/caiyun"
	"handy-translate/translate/youdao"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

type Transalte interface {
	PostQuery(value string) []string
}

func GetTransalteWay(name string) Transalte {
	switch name {
	case youdao.Way:
		return &youdao.Youdao{
			Translate: config.Translate{
				Key:    config.Data.Translate[name].Key,
				Secret: config.Data.Translate[name].Secret,
			},
		}

	case youdao.YaoDaoOnlineWay:
		return &youdao.YoudaoOnline{}

	case caiyun.Way:
		return &caiyun.Caiyun{
			Translate: config.Translate{
				Token: config.Data.Translate[name].Token,
			},
		}

	default:
		return &DefaultTransalte{}
	}
}

func (d *DefaultTransalte) PostQuery(query string) []string {
	url := "https://lingva.1link.fun/_next/data/dK-OKex2AGbb_H1CfdJwC/auto/zh/" + url.QueryEscape(query) + ".json"

	// 替换为你要请求的 URL

	// 发起 GET 请求
	response, err := http.Get(url)
	if err != nil {
		logrus.Println("HTTP 请求出错:", err)
		return nil
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Println("读取响应内容出错:", err)
		// return
	}
	var tr DefaultTransalte

	err = json.Unmarshal(body, &tr)
	if err != nil {
		logrus.WithError(err).Error("Unmarshal")
	}

	// fmt.Println("响应内容:", tr.PageProps.Translation)
	res := strings.ReplaceAll(tr.PageProps.Translation, "+", "")
	return []string{res}
}

type DefaultTransalte struct {
	PageProps PageProps `json:"pageProps"`
	NSsg      bool      `json:"__N_SSG"`
}
type Pronunciation struct {
	Translation string `json:"translation"`
}
type Info struct {
	DetectedSource    string        `json:"detectedSource"`
	Pronunciation     Pronunciation `json:"pronunciation"`
	Definitions       []interface{} `json:"definitions"`
	Examples          []interface{} `json:"examples"`
	Similar           []interface{} `json:"similar"`
	ExtraTranslations []interface{} `json:"extraTranslations"`
}
type Audio struct {
	Query       []int `json:"query"`
	Translation []int `json:"translation"`
}
type Initial struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Query  string `json:"query"`
}
type PageProps struct {
	Type        int     `json:"type"`
	Translation string  `json:"translation"`
	Info        Info    `json:"info"`
	Audio       Audio   `json:"audio"`
	Initial     Initial `json:"initial"`
}
