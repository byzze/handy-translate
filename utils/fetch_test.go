package utils

import (
	"fmt"
	"testing"
)

func TestMyFetch(t *testing.T) {
	res := MyFetch(`https://fanyi.baidu.com/langdetect`,
		map[string]interface{}{
			"method": "POST",
			"headers": map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			"body": "query=apple",
		})
	fmt.Println(res)
}
