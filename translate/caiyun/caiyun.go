package caiyun

import (
	"bytes"
	"encoding/json"
	"handy-translate/config"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// https://docs.caiyunapp.com/blog/2021/12/30/hello-world
const Way = "caiyun"

type Caiyun struct {
	config.Translate
}

type TranslationPayload struct {
	Source    []string `json:"source"`
	TransType string   `json:"trans_type"`
	RequestID string   `json:"request_id"`
	Detect    bool     `json:"detect"`
}

type TranslationResponse struct {
	Target []string `json:"target"`
}

func (c *Caiyun) PostQuery(source string) []string {
	url := "http://api.interpreter.caiyunai.com/v1/translator"

	// WARNING, this token is a test token for new developers,
	// and it should be replaced by your token
	token := c.Key

	payload := TranslationPayload{
		Source:    strings.Split(source, ","),
		TransType: "auto2zh",
		RequestID: "demo",
		Detect:    true,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logrus.Println("Error marshaling payload:", err)
		return nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logrus.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Println("Error sending request:", err)
		return nil
	}

	if resp.StatusCode != 200 {
		logrus.Println(resp)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Println("Error reading response body:", err)
		return nil
	}
	logrus.Println(string(respBody))
	var translationResponse TranslationResponse
	err = json.Unmarshal(respBody, &translationResponse)
	if err != nil {
		logrus.Println("Error unmarshaling response body:", err)
		return nil
	}

	return translationResponse.Target
}

/* func main() {
	source := []string{"Lingocloud is the best translation service.", "彩云小译は最高の翻訳サービスです"}
	target := Translate(source, "auto2zh")

	fmt.Println(target)
}
*/
