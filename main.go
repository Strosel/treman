package main

import (
	"log"
	"math/rand"
	"time"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/inkeliz/giohyperlink"
)

var (
	fontSize = unit.Sp(32)

	win *app.Window
	err error
)

func main() {
	rand.Seed(time.Now().UnixNano())

	font, err := loadFont()
	if err != nil {
		log.Fatal(err)
	}

	th := material.NewTheme(font)
	th.TextSize = fontSize

	go func() {
		win = app.NewWindow()
		if err := loop(win, th); err != nil {
			log.Fatal(err)
		}
	}()

	app.Main()
}

func loop(w *app.Window, th *material.Theme) error {
	var ops op.Ops

	screen := gameScreen(drules())
	for {
		e := <-w.Events()
		giohyperlink.ListenEvents(e)
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case key.Event:
			if e.Name == key.NameBack {
				screen = gameScreen(screen.Rules())
			}
			w.Invalidate()
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			screen = screen.Layout(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}
