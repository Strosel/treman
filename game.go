package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/image/colornames"
)

type game struct {
	dice  Roll
	rules []Rule

	rollClick *widget.Clickable
	newClick  *widget.Clickable
	ruleClick *widget.Clickable
}

func gameScreen(rules []Rule) Screen {
	return &game{
		rules:     rules,
		rollClick: new(widget.Clickable),
		newClick:  new(widget.Clickable),
		ruleClick: new(widget.Clickable),
	}
}

func (g *game) Layout(gtx Ctx, th *material.Theme) (nextScreen Screen) {
	nextScreen = g

	rolled := func(gtx Ctx) Dim {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			RigidInset(layout.Inset{Top: unit.Dp(16)}, func(gtx Ctx) Dim {
				return layout.Flex{
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					layout.Rigid(func(gtx Ctx) Dim {
						bttn := material.Button(th, g.ruleClick, "?")
						bttn.Color = colornames.Black
						bttn.Background = color.RGBA{255, 255, 255, 255}

						for g.ruleClick.Clicked() {
							nextScreen = viewRulesScreen(g.rules)
						}

						return bttn.Layout(gtx)
					}),
				)
			}),
			RigidInset(layout.UniformInset(unit.Dp(16)), func(gtx Ctx) Dim {
				dice := material.H2(th, fmt.Sprintf("%v %v", g.dice[0], g.dice[1]))
				dice.Alignment = text.Middle
				dice.Font.Variant = "Dice"
				if g.dice[0] > 6 {
					dice.Color = colornames.Rosybrown
					dice.Text = fmt.Sprintf("%v %v", (g.dice[0]%6)+1, (g.dice[1]%6)+1)
				}
				return dice.Layout(gtx)
			}),
		)
	}

	text := func(gtx Ctx) Dim {
		rolls := ""

		if g.dice[0] < 7 {
			for _, r := range g.rules {
				if r.Valid(g.dice) {
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

		lbl := material.H5(th, rolls)
		lbl.Alignment = text.Middle
		return lbl.Layout(gtx)
	}

	buttons := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(8))
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			RigidInset(in, func(gtx Ctx) Dim {
				newBttn := material.Button(th, g.newClick, "\nNy Regel\n")
				newBttn.Background = colornames.Mediumseagreen

				if (SetRule{Set: Roll{6, 6}}.Valid(g.dice)) {
					for g.newClick.Clicked() {
						nextScreen = addRuleScreen(th, g.rules)
					}
					return newBttn.Layout(gtx)
				} else if (SetRule{Set: Roll{2, 1}}.Valid(g.dice)) {
					newBttn.Text = "\nUtmaning\n"
					for g.newClick.Clicked() {
						nextScreen = challengeScreen(g.rules, g.dice)
					}
					return newBttn.Layout(gtx)
				}
				return Dim{}
			}),
			RigidInset(in, func(gtx Ctx) Dim {
				rollBttn := material.Button(th, g.rollClick, "\nRulla\n")
				for g.rollClick.Clicked() && g.dice[0] < 7 {
					go g.dice.AnimateRoll()
				}
				return rollBttn.Layout(gtx)
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

	return nextScreen
}

func (g *game) Rules() []Rule {
	return g.rules
}
