package main

import (
	"context"
	"fmt"
	"handy-translate/hook"

	"github.com/go-vgo/robotgo"
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

func (b *App) SendDataToJS(query, result string) {
	runtime.EventsEmit(b.ctx, "query", query)
	runtime.EventsEmit(b.ctx, "result", result)

}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	go hook.Hook(ctx)
	go func() {
		for {
			select {
			case <-hook.HookCenterChan:
				a.SendDataToJS("nihao", "hello 正常自行车")
				// text := hook.GetQueryText()
				// if curContent == text {
				// 	show()
				// 	continue
				// }

				// hook.SetCurText(text)
				// t.QueryContent = text

				// var transalteTool = config.Data.TranslateWay
				// way := translate.GetTransalteWay(transalteTool)
				// result := way.PostQuery(text)
				// logrus.WithField("result", result).Info("Transalte")

				// switch way.(type) {
				// case *youdao.Youdao:
				// 	t.Explain = result[0]
				// 	t.ExplainEx = result[1]
				// case *caiyun.Caiyun:
				// 	t.QueryContent = text
				// 	t.Explain = strings.Join(result, ",")
				// }

				// var b bytes.Buffer
				// tmpl.Execute(&b, t)
				// content := b.String()

				x, y := robotgo.GetMousePos()
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
