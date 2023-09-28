package translate

import (
	"context"
	"fmt"
	"handy-translate/config"
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
	config.Init(context.TODO())
	list := GetTransalteWay()
	for _, v := range list {
		s, err := v.PostQuery("tr")
		if err != nil {
			t.Fatal(err)
			continue
		}
		if len(s) == 0 {
			continue
		}
		fmt.Println(s)
		break
	}
}
