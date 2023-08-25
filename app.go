package main

import (
	"context"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate"
	"strings"

	"github.com/getlantern/systray"
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

func (a *App) SendDataToJS(query, result, explian string) {
	runtime.EventsEmit(a.ctx, "query", query)
	runtime.EventsEmit(a.ctx, "result", result)
	runtime.EventsEmit(a.ctx, "explian", explian)
}

// test data
func (a *App) onDomReady(ctx context.Context) {
	runtime.EventsEmit(a.ctx, "transalteWay", config.Data.TranslateWay)

	var translateMap = make(map[string]config.Translate)
	for k, v := range config.Data.Translate {
		tmp := config.Translate{
			Name: v.Name,
		}
		translateMap[k] = tmp
	}
	bTranslate, err := json.Marshal(translateMap)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}

	runtime.EventsEmit(a.ctx, "transalteWay", config.Data.TranslateWay)
	runtime.EventsEmit(a.ctx, "transalteMap", string(bTranslate))

	runtime.EventsEmit(a.ctx, "transalteWay", config.Data.TranslateWay)

	runtime.EventsOn(a.ctx, "transalteWay-send", func(optionalData ...interface{}) {
		logrus.WithField("optionalData", optionalData).Info("transalteWay-send")
		if len(optionalData) >= 1 {
			config.Data.TranslateWay = optionalData[0].(string)
			runtime.EventsEmit(a.ctx, "transalteWay", optionalData[0].(string))
		}
	})

	runtime.EventsOn(a.ctx, "key-save", func(optionalData ...interface{}) {
		// 遍历 optionalData 切片，对每个元素进行类型断言并转换为 string
		config.Data.Keyboard = []string{}
		for _, item := range optionalData {
			for _, items := range item.([]interface{}) {
				if str, ok := items.(string); ok {
					if strings.TrimSpace(str) == "" {
						continue
					}
					config.Data.Keyboard = append(config.Data.Keyboard, str)
				} else {
					fmt.Println("Element is not of type string:", item)
				}
			}
		}
		logrus.Info(config.Data.Keyboard)
		if len(config.Data.Keyboard) > 0 {
			runtime.EventsEmit(a.ctx, "transalteWay", config.Data.Keyboard[:len(config.Data.Keyboard)-1])
		}
		go hook.Hook(config.Data.Keyboard)
	})
	// a.SendDataToJS("Board", "董事会", "n. 板，木板；黑板，告示牌；董事会，理事会；膳食，伙食，膳食费用；局；<非正式>舞台；<美>（冰球场周围的）界墙；<旧>（美国大学的）入学考试,v. 登上（火车、轮船或飞机）；让乘客登机（或上船等）；寄宿；（在学校）住校；将（宠物）暂时寄养在他处；用木板覆盖,【名】")
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {

	go hook.Hook(config.Data.Keyboard)

	// scList, _ := runtime.ScreenGetAll(ctx)

	// var screenX, screenY int
	// for _, v := range scList {
	// 	if v.IsCurrent {
	// 		screenX = v.Width
	// 		screenY = v.Height
	// 	}
	// }

	runtime.WindowCenter(ctx)
	go func() {
		for {
			select {
			case <-hook.HookChan:
				robotgo.KeyTap("c", "ctrl")
				// windowX, windowY := runtime.WindowGetSize(ctx)
				// x, y := robotgo.GetMousePos()
				// x, y = x+10, y-10
				runtime.WindowShow(ctx)
				queryText, _ := runtime.ClipboardGetText(a.ctx)

				if queryText != hook.GetCurText() {
					hook.SetCurText(queryText)
					// 加载动画变量
					runtime.EventsEmit(a.ctx, "loading", "true")

					var transalteWay = config.Data.TranslateWay
					way := translate.GetTransalteWay(transalteWay)
					result := way.PostQuery(queryText)

					logrus.WithFields(logrus.Fields{
						"queryText":    queryText,
						"result":       result,
						"transalteWay": transalteWay,
					}).Info("Transalte")

					switch way.(type) {
					// case *youdao.Youdao:
					// 	if len(result) >= 2 {
					// 		a.SendDataToJS(queryText, result[0], result[1])
					// 	}

					default:
						a.SendDataToJS(queryText, strings.Join(result, ","), "")
					}
				}
				// TODO 弹出窗口根据鼠标位置变动
				// fmt.Println("or:", x, y, screenX, screenY, windowX, windowY)
				// if y+windowY+20 >= screenY {
				// 	y = screenY - windowY - 20
				// }

				// if x+windowX >= screenX {
				// 	x = screenX - windowX
				// }
				// fmt.Println("new:", x, y, screenX, screenY, windowX, windowY)

				// runtime.WindowSetPosition(ctx, x, y)

			}
		}
	}()
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
	systray.Quit()
}
