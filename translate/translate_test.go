package translate

import (
	"fmt"
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
