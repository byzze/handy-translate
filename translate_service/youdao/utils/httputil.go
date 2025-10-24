package utils

import (
	"io"
	"log/slog"
	"net/http"
	neturl "net/url"
	"strings"
	"time"
)

func DoGet(url string, header map[string][]string, paramsMap map[string][]string, expectContentType string) []byte {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	params := neturl.Values{}
	for k, v := range paramsMap {
		params[k] = v
	}
	parseUrl, _ := neturl.Parse(url)
	parseUrl.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", parseUrl.String(), nil)
	for k, v := range header {
		for hv := range v {
			req.Header.Add(k, v[hv])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		slog.Error("request failed:", slog.Any("err", err))
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, expectContentType) {
		slog.Error("contentType not match", slog.String("contentType", contentType), slog.String("expectContentType", expectContentType))
		return nil
	}
	return body
}

func DoPost(url string, header map[string][]string, bodyMap map[string][]string, expectContentType string) []byte {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	params := neturl.Values{}
	for k, v := range bodyMap {
		for pv := range v {
			params.Add(k, v[pv])
		}
	}
	req, _ := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	for k, v := range header {
		for hv := range v {
			req.Header.Add(k, v[hv])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		slog.Error("request failed:", slog.Any("err", err))
		return nil
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, expectContentType) {
		slog.Error("contentType not match", slog.String("contentType", contentType), slog.String("expectContentType", expectContentType))
		return nil
	}
	return body
}
