package baidu

import (
	"fmt"
	"handy-translate/config"
	"testing"
)

func TestBaidu_PostQuery(t *testing.T) {
	source := `hello`
	var baidu = &Baidu{
		Translate: config.Translate{
			Key: "A0BWMb27a5cADOXeTKio",
		},
	}
	target := baidu.PostQuery(source)

	fmt.Println(target)
}
