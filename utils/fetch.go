package utils

import (
	"bufio"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

// MyFetch 封装请求，因为从前端发起http请求会出现跨域
func MyFetch(URL string, content map[string]interface{}) interface{} {
	client := &http.Client{}
	var req *http.Request
	var err error

	var method = http.MethodGet
	if v, ok := content["method"]; ok {
		method = fmt.Sprintf("%v", v)
	}

	bodycontetn := content["body"]
	body := fmt.Sprintf("%v", bodycontetn)

	if method == "GET" {
		req, err = http.NewRequest(method, URL+"?"+body, nil)
	} else {
		req, err = http.NewRequest(method, URL, strings.NewReader(body))
	}

	if err != nil {
		slog.Error("err", slog.Any("err", err))
		return err
	}

	if h, ok := content["headers"]; ok {
		if hMap, hcok := h.(map[string]interface{}); hcok {
			for k, val := range hMap {
				valstr := fmt.Sprintf("%v", val)
				req.Header.Set(k, valstr)
			}
		}
	}

	slog.Info("req", slog.Any("req", req))

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(resp.Body)

	var responseBody string
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		responseBody += trimmedLine
	}

	if err := scanner.Err(); err != nil {
		slog.Error("err", slog.Any("err", err))
		return err
	}

	return responseBody
}
