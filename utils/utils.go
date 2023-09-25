package utils

import (
	"fmt"
	"unicode"
)

func CountQueryStr(str string) (chineseCount, englishCount int) {
	for _, char := range str {
		if unicode.Is(unicode.Han, char) || unicode.Is(unicode.Hangul, char) {
			// 判断字符是否为中文（汉字）或韩文
			chineseCount++
		} else if unicode.Is(unicode.Latin, char) {
			// 判断字符是否为拉丁字母（英文字符）
			englishCount++
		}
	}
	fmt.Printf("英文字母数: %d\n", englishCount)
	fmt.Printf("中文字数: %d\n", chineseCount)
	return
}
