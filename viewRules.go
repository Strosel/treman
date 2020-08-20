package main

import (
	"image/color"
	"runtime"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/image/colornames"
)

var baserules = `Först väljs en spelare till treman, detta är en titel som kommer förflytta sig under spelets gång. 

Titeln treman förflyttas till en ny spelare vid ett av följande tillfällen:
• Om en person går med i spelet efter start är denne nu treman.
• Om en splare lämnar bordet och kommer tillbaks (oftast för att hämta mer dricka) är denne nu treman.
• Vid vissa tärningsslag blir antingen den som slagit tärningarna eller en annan spelare ny treman. (se nedan)
När en ny person blir treman skålar denne med gammla treman och bägge dricker.

När treman valts börjar en spelare slå tärningarna och följer instruktionerna som tillhör vad de slagit (se nedan). Samma spelare fortsätter slå tärningarna tills de får "Ingenting", dvs. ett slag utan tillhörande regel, och då skickas tärningarna vidare medsols.

n+n klunkar betyder att n st klunkar får delas ut till två personer eller dubbelt så mycket till en person`

type viewRules struct {
	rules []Rule

	list        *layout.List
	cancelClick *widget.Clickable
}

func viewRulesScreen(rules []Rule) Screen {
	return &viewRules{
		rules: rules,
		list: &layout.List{
			Axis: layout.Vertical,
		},
		cancelClick: new(widget.Clickable),
	}
}

func (v *viewRules) Layout(gtx Ctx, th *material.Theme) (nextScreen Screen) {
	nextScreen = v

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx Ctx) Dim {
		return v.list.Layout(gtx, len(v.rules)+2, func(gtx Ctx, i int) Dim {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx Ctx) Dim {
				if i == 0 && runtime.GOOS == "android" {
					return layout.Inset{Top: unit.Dp(16)}.Layout(gtx, material.H6(th, "Regler").Layout)
				} else if i == 0 {
					bttn := material.Button(th, v.cancelClick, "← Regler"+runtime.GOOS)
					bttn.Color = colornames.Black
					bttn.Background = color.RGBA{255, 255, 255, 255}
					bttn.TextSize = material.H6(th, "").TextSize
					bttn.Inset.Left = unit.Dp(0)

					for v.cancelClick.Clicked() {
						nextScreen = gameScreen(v.rules)
					}

					return layout.Flex{
						Alignment: layout.Start,
					}.Layout(gtx,
						RigidInset(layout.Inset{Top: unit.Dp(16)}, bttn.Layout),
					)
				} else if i == 1 {
					body := material.Body1(th, baserules)
					body.TextSize = unit.Dp(24)
					return body.Layout(gtx)
				}
				return v.rules[i-2].Widget(th)(gtx)
			})
		})
	})

	return nextScreen
}

func (v *viewRules) Rules() []Rule {
	return v.rules
}
