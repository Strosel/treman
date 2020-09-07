package main

import (
	"fmt"
	"image/color"
	"math"

	"gioui.org/layout"
	"gioui.org/text"
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

func DiceLayout(th *material.Theme, d int, c ...color.RGBA) func(Ctx) Dim {
	if len(c) == 0 {
		c = []color.RGBA{colornames.Black}
	}

	dice := material.H2(th, fmt.Sprint(d))
	dice.Alignment = text.Middle
	dice.Font.Variant = "Dice"
	dice.Color = c[0]
	if d > 6 {
		dice.Color = c[len(c)-1]
		dice.Text = fmt.Sprint((d % 6) + 1)
	}
	return dice.Layout
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
