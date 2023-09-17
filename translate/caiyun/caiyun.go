package caiyun

import (
	"bytes"
	"encoding/json"
	"handy-translate/config"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// https://docs.caiyunapp.com/blog/2021/12/30/hello-world
const Way = "彩云翻译"

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

func (c *Caiyun) GetName() string {
	return Way
}

func (c *Caiyun) PostQuery(source string) ([]string, error) {
	url := "http://api.interpreter.caiyunai.com/v1/translator"

	// WARNING, this token is a test token for new developers,
	// and it should be replaced by your token
	token := c.Token

	payload := TranslationPayload{
		Source:    strings.Split(source, ","),
		TransType: "auto2zh",
		RequestID: "demo",
		Detect:    true,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logrus.Println("Error marshaling payload:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logrus.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Println("Error sending request:", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		logrus.Println(resp)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Println("Error reading response body:", err)
		return nil, err
	}
	logrus.Println(string(respBody))
	var translationResponse TranslationResponse
	err = json.Unmarshal(respBody, &translationResponse)
	if err != nil {
		logrus.Println("Error unmarshaling response body:", err)
		return nil, err
	}

	return translationResponse.Target, nil
}

/* func main() {
	source := []string{"Lingocloud is the best translation service.", "彩云小译は最高の翻訳サービスです"}
	target := Translate(source, "auto2zh")

	fmt.Println(target)
}
*/
