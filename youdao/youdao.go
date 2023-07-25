package youdao

import (
	"encoding/json"
	"strings"
	"translate/youdao/utils"
	"translate/youdao/utils/authv3"

	"github.com/sirupsen/logrus"
)

// 您的应用ID
var appKey = "5ade39ff86f06b4a"

// 您的应用密钥
var appSecret = "UCAXNIgHOebOlxosl67ZLf2wXQnuhrAJ"

func PostQuery(query string) []string {
	// 添加请求参数
	paramsMap := createRequestParams(query)
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/api", header, paramsMap, "application/json")
	// 打印返回结果

	var tr Transalte

	err := json.Unmarshal(result, &tr)
	if err != nil {
		logrus.Println(err)
		return nil
	}
	transalteResult := strings.Join(tr.Translation, ",")
	transalteExplains := strings.Join(tr.Basic.Explains, ",")

	return []string{transalteResult, transalteExplains}
}

func createRequestParams(query string) map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91/API%E6%96%87%E6%A1%A3/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1-API%E6%96%87%E6%A1%A3.html
	*/
	q := query
	from := "auto"
	to := "zh-CHS"
	vocabId := "您的用户词表ID"

	return map[string][]string{
		"q":       {q},
		"from":    {from},
		"to":      {to},
		"vocabId": {vocabId},
	}
}

type Transalte struct {
	ReturnPhrase  []string      `json:"returnPhrase"`
	Query         string        `json:"query"`
	ErrorCode     string        `json:"errorCode"`
	L             string        `json:"l"`
	TSpeakURL     string        `json:"tSpeakUrl"`
	Web           []Web         `json:"web"`
	RequestID     string        `json:"requestId"`
	Translation   []string      `json:"translation"`
	MTerminalDict MTerminalDict `json:"mTerminalDict"`
	Dict          Dict          `json:"dict"`
	Webdict       Webdict       `json:"webdict"`
	Basic         Basic         `json:"basic"`
	IsWord        bool          `json:"isWord"`
	SpeakURL      string        `json:"speakUrl"`
}
type Web struct {
	Value []string `json:"value"`
	Key   string   `json:"key"`
}
type MTerminalDict struct {
	URL string `json:"url"`
}
type Dict struct {
	URL string `json:"url"`
}
type Webdict struct {
	URL string `json:"url"`
}
type Wf struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Wfs struct {
	Wf Wf `json:"wf"`
}
type Basic struct {
	ExamType   []string `json:"exam_type"`
	UsPhonetic string   `json:"us-phonetic"`
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	Wfs        []Wfs    `json:"wfs"`
	UkSpeech   string   `json:"uk-speech"`
	Explains   []string `json:"explains"`
	UsSpeech   string   `json:"us-speech"`
}
