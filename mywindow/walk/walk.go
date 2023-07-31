package walk

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Run() {
	var mainWindow *walk.MainWindow

	MainWindow{
		Title:    "My Window",
		Size:     Size{Width: 400, Height: 300},
		Layout:   VBox{},
		AssignTo: &mainWindow,
	}.Run()

	// 设置窗口位置
	if mainWindow != nil {
		go func() {
			for {
				time.Sleep(time.Second * 3)
				x, y := robotgo.GetMousePos()
				mainWindow.SetBoundsPixels(walk.Rectangle{X: x, Y: y, Width: 400, Height: 300})
			}
		}()
	}
}
