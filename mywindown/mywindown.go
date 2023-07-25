package mywindown

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
)

var MyWindown fyne.Window

var myLk sync.Mutex

// Show 处理数据竞态
func Show() {
	fmt.Println("a ...any")
	if myLk.TryLock() {
		fmt.Println("111")
		MyWindown.Show()
		myLk.Unlock()
		fmt.Println("222")
	}
}
