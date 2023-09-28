package baidu

import (
	"context"
	"fmt"
	"handy-translate/config"
	"testing"
)

func TestBaidu_PostQuery(t *testing.T) {
	config.Init(context.TODO())
	source := `Number of English letters`
	var baidu = &Baidu{
		Translate: config.Translate{
			Key:   config.Data.Translate[0].Key,
			AppID: config.Data.Translate[0].AppID,
		},
	}
	target, err := baidu.PostQuery(source)
	fmt.Println(err)
	fmt.Println(target)
}
