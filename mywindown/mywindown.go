package mywindown

import (
	"sync"

	"fyne.io/fyne/v2"
)

var MyWindown fyne.Window

var myLk sync.Mutex

// Show 处理数据竞态
func Show() {
	myLk.Lock()
	MyWindown.Show()
	myLk.Unlock()
}
