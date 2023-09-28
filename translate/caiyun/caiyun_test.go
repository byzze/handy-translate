package caiyun

import (
	"fmt"
	"handy-translate/config"
	"testing"
)

func TestTranslate(t *testing.T) {
	// source := []string{"Lingocloud is the best translation service.", "彩云小译は最高の翻訳サービスです"}
	// target := Translate(source, "auto2zh")

	// fmt.Println(target)
	source := `hello`
	var caiyun = &Caiyun{
		Translate: config.Translate{
			Token: "9t86wdbb14mx8o9qhouq",
		},
	}
	target, _ := caiyun.PostQuery(source)

	fmt.Println(target)
}
