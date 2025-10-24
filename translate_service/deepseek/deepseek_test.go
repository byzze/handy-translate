package deepseek

import (
	"fmt"
	"handy-translate/config"
	"testing"
)

func TestDeepseek_PostQuery(t *testing.T) {
	config.Init("handy-translate")
	source := `hello`
	var deepseek = &Deepseek{
		Translate: config.Translate{
			Key:   config.Data.Translate[Way].Key,
			AppID: config.Data.Translate[Way].AppID,
		},
	}

	target, err := deepseek.PostQuery(source, "auto", "en")
	fmt.Println(err)
	fmt.Println(target)
}
