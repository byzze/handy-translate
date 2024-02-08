package translate_service

import (
	"fmt"
	"handy-translate/config"
	"handy-translate/translate_service/baidu"
	"handy-translate/translate_service/youdao"
	"testing"

	"github.com/OwO-Network/gdeeplx"
)

func TestGetTransalteWay(t *testing.T) {
	result, err := gdeeplx.Translate("hello", "EN", "ZH", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)
}

func TestGetTransalteWayList(t *testing.T) {
	config.Init("handy-translate")
	v := GetTransalteWay(baidu.Way)
	s, err := v.PostQuery("app", "auto", "zh")
	if err != nil {
		t.Fatal(err)

	}
	fmt.Println(s)
}

func TestTranslateYouDao(t *testing.T) {
	config.Init("handy-translate")
	v := GetTransalteWay(youdao.Way)
	s, err := v.PostQuery("china", "auto", "zh")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
}
