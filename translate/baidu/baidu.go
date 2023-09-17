package baidu

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"io"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

// https://docs.caiyunapp.com/blog/2021/12/30/hello-world
const Way = "百度翻译"

type Baidu struct {
	config.Translate
}

const (
	appID = "20230823001790949"

	fromLang = "en"
	toLang   = "zh"
	endpoint = "http://api.fanyi.baidu.com"
	path     = "/api/trans/vip/translate"
)

func (b *Baidu) GetName() string {
	return Way
}

func (b *Baidu) PostQuery(source string) ([]string, error) {
	appKey := b.Key

	query := source

	salt := rand.Intn(32768) + 32768
	sign := generateMD5(appID + query + fmt.Sprint(salt) + appKey)

	data := url.Values{}
	data.Add("appid", appID)
	data.Add("q", query)
	data.Add("from", fromLang)
	data.Add("to", toLang)
	data.Add("salt", fmt.Sprint(salt))
	data.Add("sign", sign)

	resp, err := http.PostForm(endpoint+path, data)
	if err != nil {
		logrus.Error("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Error:", err)
		return nil, err
	}

	var result translateResult
	if err := json.Unmarshal(body, &result); err != nil {
		logrus.Error("Error:", err)
		return nil, err
	}

	prettyResult, _ := json.MarshalIndent(result, "", "    ")
	logrus.Println(string(prettyResult))

	if len(result.TransResult) > 0 {
		if result.TransResult[0].Dst == result.TransResult[0].Src {
			return nil, nil
		}
		return []string{result.TransResult[0].Dst}, nil
	}
	return nil, err
}

func generateMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", hash)
}

type translateResult struct {
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
}
type TransResult struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}
