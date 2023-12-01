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
