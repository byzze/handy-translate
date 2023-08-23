package caiyun

import (
	"fmt"
	"handy-translate/config"
	"testing"
)

func TestTranslate(t *testing.T) {
	source := `hello`
	var caiyun = &Caiyun{
		Translate: config.Translate{
			Token: "9t86wdbb14mx8o9qhouq",
		},
	}
	target := caiyun.PostQuery(source)

	fmt.Println(target)
}
