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

type game struct {
	dice  Roll
	rules []Rule

	rollClick *widget.Clickable
	newClick  *widget.Clickable
}

func gameScreen(rules []Rule) Screen {
	return &game{
		rules:     rules,
		rollClick: new(widget.Clickable),
		newClick:  new(widget.Clickable),
	}
}

func (g *game) Layout(gtx Ctx, th *material.Theme) (nextScreen Screen) {
	nextScreen = g

	rolled := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(16))
		in.Top = unit.Dp(64)
		return layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			RigidInset(in, widget.Image{Src: sprites[g.dice[0]], Scale: 2}.Layout),
			RigidInset(in, widget.Image{Src: sprites[g.dice[1]], Scale: 2}.Layout),
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
			RigidInset(in, func(gtx Ctx) Dim {
				newBttn := material.Button(th, g.newClick, "\nNy Regel\n")
				newBttn.Background = colornames.Mediumseagreen

				if (SetRule{Set: Roll{6, 6}}.Valid(g.dice)) {
					for g.newClick.Clicked() {
						nextScreen = addRuleScreen(g.rules)
					}
					return newBttn.Layout(gtx)
				}
				return Dim{}
			}),
			RigidInset(in, func(gtx Ctx) Dim {
				rollBttn := material.Button(th, g.rollClick, "\nRulla\n")
				for g.rollClick.Clicked() {
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
