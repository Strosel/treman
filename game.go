package main

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/image/colornames"
)

var (
	rollButton = new(widget.Clickable)
	newButton  = new(widget.Clickable)
)

func game(gtx Ctx, th *material.Theme) {

	rolled := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(16))
		in.Top = unit.Dp(64)
		return layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			layout.Rigid(func(gtx Ctx) Dim {
				return in.Layout(gtx, widget.Image{Src: sprites[dice[0]], Scale: 2}.Layout)
			}),
			layout.Rigid(func(gtx Ctx) Dim {
				return in.Layout(gtx, widget.Image{Src: sprites[dice[1]], Scale: 2}.Layout)
			}),
		)
	}

	text := func(gtx Ctx) Dim {
		rolls := ""

		if !rolling {
			for _, r := range rules {
				if r.Valid(dice) {
					if len(rolls) == 0 {
						rolls += r.String()
					} else {
						rolls += fmt.Sprintf(", %v", r.String())
					}
				}
			}

			if len(rolls) == 0 {
				rolls = "Ingenting"
			}
		}

		lbl := material.Label(th, bigFont, rolls)
		lbl.Alignment = text.Middle
		return lbl.Layout(gtx)
	}

	buttons := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(8))
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			layout.Rigid(func(gtx Ctx) Dim {
				return in.Layout(gtx, func(gtx Ctx) Dim {
					if (SetRule{Set: Roll{6, 6}}.Valid(dice)) {
						bttn := material.Button(th, newButton, "\nNy Regel\n")
						bttn.Background = colornames.Mediumseagreen
						for newButton.Clicked() {
							playing = false
							ruleRadio.Value = "sum"
						}
						return bttn.Layout(gtx)
					}
					return Dim{}
				})
			}),
			layout.Rigid(func(gtx Ctx) Dim {
				return in.Layout(gtx, func(gtx Ctx) Dim {
					bttn := material.Button(th, rollButton, "\nRulla\n")
					for rollButton.Clicked() {
						go dice.AnimateRoll()
					}
					return bttn.Layout(gtx)
				})
			}),
		)
	}

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx Ctx) Dim {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			layout.Rigid(rolled),
			layout.Flexed(1, text),
			layout.Rigid(buttons),
		)
	})
}
