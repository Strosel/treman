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
	rollButton = new(widget.Button)
	newButton  = new(widget.Button)
)

func game(gtx *layout.Context, th *material.Theme) {

	paint := func() {
		in := layout.UniformInset(unit.Dp(16))
		in.Top = unit.Dp(64)
		layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			layout.Rigid(func() {
				in.Layout(gtx, func() {
					th.Image(sprites[dice[0]]).Layout(gtx)
				})
			}),
			layout.Rigid(func() {
				in.Layout(gtx, func() {
					th.Image(sprites[dice[1]]).Layout(gtx)
				})
			}),
		)
	}

	text := func() {
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

		lbl := th.Label(bigFont, rolls)
		lbl.Alignment = text.Middle
		lbl.Layout(gtx)
	}

	buttons := func() {
		in := layout.UniformInset(unit.Dp(8))
		layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			layout.Rigid(func() {
				in.Layout(gtx, func() {
					if (SetRule{Set: Roll{6, 6}}.Valid(dice)) {
						bttn := th.Button("\nNy Regel\n")
						bttn.TextSize = fontSize
						bttn.Background = colornames.Mediumseagreen
						for newButton.Clicked(gtx) {
							playing = false
							ruleRadio.SetValue("sum")
						}
						bttn.Layout(gtx, newButton)
					}
				})
			}),
			layout.Rigid(func() {
				in.Layout(gtx, func() {
					bttn := th.Button("\nRulla\n")
					bttn.TextSize = fontSize
					for rollButton.Clicked(gtx) {
						go dice.AnimateRoll()
					}
					bttn.Layout(gtx, rollButton)
				})
			}),
		)
	}

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func() {
		layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			layout.Rigid(paint),
			layout.Flexed(1, text),
			layout.Rigid(buttons),
		)
	})
}
