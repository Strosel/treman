package main

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type challenge struct {
	dice  Roll
	rules []Rule

	bttn *widget.Clickable
}

func challengeScreen(rules []Rule) Screen {
	return &challenge{
		rules: rules,
		bttn:  new(widget.Clickable),
	}
}

func (c *challenge) Rules() []Rule {
	return c.rules
}

func (c *challenge) Layout(gtx Ctx, th *material.Theme) (nextScreen Screen) {
	nextScreen = c

	title := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(0))
		in.Top = unit.Dp(32)
		lbl := material.H4(th, "Utmaning")
		lbl.Alignment = text.Middle
		return in.Layout(gtx, lbl.Layout)
	}

	dice := func(gtx Ctx) Dim {
		return layout.Flex{
			Spacing: layout.SpaceEvenly,
		}.Layout(gtx,
			layout.Rigid(DiceLayout(th, c.dice[0], MYRED)),
			layout.Rigid(DiceLayout(th, c.dice[1], ROYALBLUE)),
		)
	}

	text := func(gtx Ctx) Dim {
		lbl := material.H5(th, "Välj vars en tärning")
		if c.dice[0] > 6 {
			lbl.Text = ""
		} else if c.dice[0] > c.dice[1] {
			lbl.Text = "Röd är ny treman"
		} else if c.dice[0] < c.dice[1] {
			lbl.Text = "Blå är ny treman"
		}
		lbl.Alignment = text.Middle
		return lbl.Layout(gtx)
	}

	button := func(gtx Ctx) Dim {
		return layout.UniformInset(unit.Dp(8)).Layout(gtx,
			func(gtx Ctx) Dim {
				bttn := material.Button(th, c.bttn, "\nKör\n")
				if c.dice[0] != c.dice[1] && c.dice[0] < 7 {
					bttn.Text = "\nOK\n"
				}
				bttn.Background = MEDIUMSEAGREEN

				for c.bttn.Clicked() && c.dice[0] < 7 {
					if bttn.Text == "\nKör\n" {
						go func() {
							c.dice.AnimateRoll()
							for c.dice[0] == c.dice[1] {
								//Challange should never yield identical dice
								c.dice.Roll()
							}
							win.Invalidate()
						}()
					} else {
						nextScreen = gameScreen(c.rules)
					}
				}

				return bttn.Layout(gtx)
			})
	}

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx Ctx) Dim {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
			Spacing:   layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(title),
			layout.Rigid(dice),
			layout.Rigid(text),
			layout.Rigid(button),
		)
	})

	if nextScreen, ok := nextScreen.(*game); ok {
		go nextScreen.dice.AnimateRoll()
	}

	return nextScreen
}
