package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var baserules = `Först väljs en spelare till treman, detta är en titel som kommer förflytta sig under spelets gång. 

Titeln treman förflyttas till en ny spelare vid ett av följande tillfällen:
• Om en person går med i spelet efter start är denne nu treman.
• Om en splare lämnar bordet och kommer tillbaks (oftast för att hämta mer dricka) är denne nu treman.
• Vid vissa tärningsslag blir antingen den som slagit tärningarna eller en annan spelare ny treman. (se nedan)
När en ny person blir treman skålar denne med gammla treman och bägge dricker.

När treman valts börjar en spelare slå tärningarna och följer instruktionerna som tillhör vad de slagit (se nedan). Samma spelare fortsätter slå tärningarna tills de får "Ingenting", dvs. ett slag utan tillhörande regel, och då skickas tärningarna vidare medsols.

Formatet n+n klunkar betyder att n st klunkar får delas ut till två personer eller dubbelt så mycket till en person`

type viewRules struct {
	rules []Rule

	list *layout.List
}

func viewRulesScreen(rules []Rule) Screen {
	return &viewRules{
		rules: rules,
		list: &layout.List{
			Axis: layout.Vertical,
		},
	}
}

func (v *viewRules) Layout(gtx Ctx, th *material.Theme) (nextScreen Screen) {
	nextScreen = v

	th.TextSize = unit.Dp(24)
	layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx Ctx) Dim {
		return v.list.Layout(gtx, len(v.rules)+2, func(gtx Ctx, i int) Dim {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx Ctx) Dim {
				if i == 0 {
					return material.H4(th, "Regler").Layout(gtx)
				} else if i == 1 {
					return material.Body1(th, baserules).Layout(gtx)
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
