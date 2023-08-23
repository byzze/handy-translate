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
)

// https://docs.caiyunapp.com/blog/2021/12/30/hello-world
const Way = "baidu"

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

func (b *Baidu) PostQuery(source string) []string {
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
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var result translateResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	prettyResult, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(prettyResult))
	if len(result.TransResult) > 0 {
		return []string{result.TransResult[0].Dst}
	}
	return nil
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
