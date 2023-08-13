package main

import (
	"context"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate"
	"handy-translate/translate/caiyun"
	"handy-translate/translate/youdao"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (b *App) SendDataToJS(query, result, explian string) {
	runtime.EventsEmit(b.ctx, "query", query)
	runtime.EventsEmit(b.ctx, "result", result)
	runtime.EventsEmit(b.ctx, "explian", explian)
}

func (a *App) onDomReady(ctx context.Context) {
	a.SendDataToJS("Board", "董事会", "n. 板，木板；黑板，告示牌；董事会，理事会；膳食，伙食，膳食费用；局；<非正式>舞台；<美>（冰球场周围的）界墙；<旧>（美国大学的）入学考试,v. 登上（火车、轮船或飞机）；让乘客登机（或上船等）；寄宿；（在学校）住校；将（宠物）暂时寄养在他处；用木板覆盖,【名】 （Board）（英、西）博德（人名）The onMount function tells Svelte compiler to run the code while mounting this component. This code initializes the theme stored in the $theme state store. The function then sets up the event receiver to receive “message” events using the rt.EventsOn function. When it gets a message, it sets the $message store variable to it and sets the $state store variable to “message”. This tells the HTML code to show the Message component and show the window using rt.ShowWindow().")
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	go hook.Hook(ctx)

	scList, _ := runtime.ScreenGetAll(ctx)

	var screenX, screenY int
	for _, v := range scList {
		if v.IsCurrent {
			screenX = v.Width
			screenY = v.Height
		}
	}

	go func() {
		for {
			select {
			case <-hook.HookCenterChan:
				windowX, windowY := runtime.WindowGetSize(ctx)
				x, y := robotgo.GetMousePos()

				queryText, _ := runtime.ClipboardGetText(a.ctx)

				if queryText != hook.GetCurText() {
					hook.SetCurText(queryText)

					var transalteTool = config.Data.TranslateWay
					way := translate.GetTransalteWay(transalteTool)
					result := way.PostQuery(queryText)
					logrus.WithField("result", result).Info("Transalte")

					switch way.(type) {
					case *youdao.Youdao:
						if len(result) >= 2 {
							a.SendDataToJS(queryText, result[0], result[1])
						}

					case *caiyun.Caiyun:
						a.SendDataToJS(queryText, strings.Join(result, ","), "")

					default:
						a.SendDataToJS(queryText, "translate failed", "")
					}
				}
				fmt.Println("or:", x, y, screenX, screenY, windowX, windowY)

				if y+windowY+200 >= screenY {
					y = screenY - 2*windowY - 200
				}

				if x+windowX >= screenX {
					x = screenX - 2*windowX
				}
				fmt.Println("new:", x, y, screenX, screenY, windowX, windowY)
				runtime.WindowSetPosition(ctx, x, y)
				runtime.WindowShow(ctx)
			}
		}
	}()
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
