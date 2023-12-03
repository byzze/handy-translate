package baidu

import (
	"fmt"
	"handy-translate/config"
	"testing"
)

func TestBaidu_PostQuery(t *testing.T) {
	config.Init("handy-translate")
	source := `Number of English letters`
	var baidu = &Baidu{
		Translate: config.Translate{
			Key:   config.Data.Translate[Way].Key,
			AppID: config.Data.Translate[Way].AppID,
		},
	}
	target, err := baidu.PostQuery(source, "", "")
	fmt.Println(err)
	fmt.Println(target)
}
