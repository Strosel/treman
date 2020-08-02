//go:generate pkger
package main

import (
	"log"
	"math/rand"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var (
	fontSize = unit.Dp(32)

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
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case *system.CommandEvent:
			if e.Type == system.CommandBack {
				screen = gameScreen(screen.Rules())
			}
			e.Cancel = true
			w.Invalidate()
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			screen = screen.Layout(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}
