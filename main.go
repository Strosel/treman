//go:generate pkger
package main

import (
	"log"
	"math/rand"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

//Ctx is a helper alias for less wide code
type Ctx = layout.Context

//Dim is a helper alias for less wide code
type Dim = layout.Dimensions

var (
	fontSize = unit.Dp(32)
	bigFont  = unit.Dp(45)
	rolling  = false

	win     *app.Window
	sprites []paint.ImageOp
	err     error
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sprites, err = loadSprites()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		win = app.NewWindow()
		if err := loop(win); err != nil {
			log.Fatal(err)
		}
	}()

	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	th.TextSize = fontSize

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
