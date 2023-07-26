package gio

import (
	"fmt"
	"image/color"
	"log"
	"lyzee-translate/register"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/sirupsen/logrus"
)

var gioWindow *app.Window

func Init() {
	go func() {
		gioWindow = app.NewWindow()
		err := run()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	go func() {
		handleHookData()
	}()
	app.Main()
}

func run() error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-gioWindow.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			fmt.Println("eerrrrrrrrrrrrrrr", e)
			return e.Err
		case system.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := layout.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(th, "Hello, Gio")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
			//
			// gtx := layout.NewContext(&ops, e)

			// title := material.H1(th, "Hello, Gio")
			// maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			// title.Color = maroon
			// title.Alignment = text.Middle
			// title.Layout(gtx)

			// e.Frame(gtx.Ops)
		// case system.StageEvent:
		// 	if e.Stage.String() == "StagePaused"
		default:
			fmt.Println("event gio", e)
		}
	}
}

func handleHookData() {
	for {
		select {
		case ev := <-register.HookChan:
			logrus.Info("handleHookData: ", ev)
			curContent := register.GetCurContent()
			queryContent := register.GetQueryContent()
			gioWindow.Option(app.Windowed.Option())
			continue
			if curContent == queryContent {
				gioWindow.Option(app.Windowed.Option())
				continue
			}

			gioWindow.Option(app.Windowed.Option())
		}
	}
}
