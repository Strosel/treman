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
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	fontSize = unit.Dp(32)
	bigFont  = unit.Dp(45)
	noButton = new(widget.Button)
	playing  = true
	rolling  = false

	win     *app.Window
	rules   = drules()
	dice    Roll
	sprites []paint.ImageOp
	err     error
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sprites, err = loadSprites()
	if err != nil {
		log.Fatal(err)
	}

	gofont.Register()
	d1Edit.SingleLine = true
	d2Edit.SingleLine = true
	nameEdit.Submit = true

	go func() {
		win = app.NewWindow()
		if err := loop(win); err != nil {
			log.Fatal(err)
		}
	}()

	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	th.TextSize = fontSize
	gtx := layout.NewContext(w.Queue())
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case *system.CommandEvent:
			if e.Type == system.CommandBack {
				playing = true
			}
			e.Cancel = true
			w.Invalidate()
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			if playing {
				game(gtx, th)
			} else {
				addRule(gtx, th)
			}
			e.Frame(gtx.Ops)
		}
	}
}
