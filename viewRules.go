package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

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

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx Ctx) Dim {
		return v.list.Layout(gtx, len(v.rules)+1, func(gtx Ctx, i int) Dim {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx Ctx) Dim {
				if i == 0 {
					return material.Body1(th, "").Layout(gtx)
				}
				return v.rules[i-1].Widget(th)(gtx)
			})
		})
	})

	return nextScreen
}

func (v *viewRules) Rules() []Rule {
	return v.rules
}
