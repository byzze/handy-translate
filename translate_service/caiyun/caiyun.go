package caiyun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"io"
	"log/slog"
	"net/http"
	"strings"
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

func (c *Caiyun) GetName() string {
	return Way
}

func (c *Caiyun) PostQuery(query, fromLang, toLang string) ([]string, error) {
	url := "http://api.interpreter.caiyunai.com/v1/translator"

	// WARNING, this token is a test token for new developers,
	// and it should be replaced by your token
	token := c.Key

	transType := fmt.Sprintf("%s2%s", fromLang, toLang)
	payload := TranslationPayload{
		Source: strings.Split(query, ","),
		// TransType: "auto2zh",
		TransType: transType,
		RequestID: "demo",
		Detect:    true,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	slog.Info(string(respBody))
	var translationResponse TranslationResponse
	err = json.Unmarshal(respBody, &translationResponse)

	if err != nil {
		return nil, err
	}

	return translationResponse.Target, nil
}
