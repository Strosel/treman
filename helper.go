package main

import (
	"fmt"
	"image/color"
	"math"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/image/colornames"
)

//Ctx is a helper alias for less wide code
type Ctx = layout.Context

//Dim is a helper alias for less wide code
type Dim = layout.Dimensions

func RigidInset(in layout.Inset, widget layout.Widget) layout.FlexChild {
	return layout.Rigid(func(gtx Ctx) Dim {
		return in.Layout(gtx, widget)
	})
}

func FlexedInset(in layout.Inset, weight float32, widget layout.Widget) layout.FlexChild {
	return layout.Flexed(weight, func(gtx Ctx) Dim {
		return in.Layout(gtx, widget)
	})
}

type diceButton struct {
	th    *material.Theme
	click *widget.Clickable
	val   int
}

func newDiceButton(th *material.Theme) *diceButton {
	return &diceButton{
		th:    th,
		click: new(widget.Clickable),
		val:   1,
	}
}

func (d *diceButton) Layout(gtx Ctx, min, max int) Dim {
	bttn := material.Button(d.th, d.click, fmt.Sprint(d.val))
	bttn.TextSize = material.H4(d.th, "").TextSize
	bttn.Color = colornames.Black
	bttn.Background = color.RGBA{255, 255, 255, 255}
	if max < 7 {
		bttn.Font.Variant = "Dice"
	}

	for d.click.Clicked() {
		d.val++
	}

	d.val = int(math.Max(float64(min), float64(d.val%(max+1))))
	return bttn.Layout(gtx)
}
