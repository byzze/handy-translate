package main

import (
	_ "embed"
	"testing"

	"github.com/lxn/win"
)

func TestShowForWindows(t *testing.T) {
	t1 := win.TRACKMOUSEEVENT{}
	win.TrackMouseEvent(&t1)
	select {}
}

func TestCaptureSelectedScreen(t *testing.T) {
	a := App{}
	a.CaptureSelectedScreen(30, 30, 300, 300)
}
