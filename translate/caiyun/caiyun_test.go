package caiyun

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
)

func TestTranslate(t *testing.T) {
	source := `Lingocloud is the best translation service.`
	var caiyun = &Caiyun{}
	target := caiyun.PostQuery(source)

	fmt.Println(target)

	x, y := robotgo.GetMousePos()
	color := robotgo.GetPixelColor(x, y)
	fmt.Println(color)
}
