package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type addRule struct {
	rules []Rule

	ruleRadio *widget.Enum
	d1        *diceButton
	d2        *diceButton
	nameEdit  *widget.Editor
	saveClick *widget.Clickable
}

func addRuleScreen(th *material.Theme, rules []Rule) Screen {
	a := &addRule{
		rules:     rules,
		ruleRadio: new(widget.Enum),
		d1:        newDiceButton(th),
		d2:        newDiceButton(th),
		nameEdit:  new(widget.Editor),
		saveClick: new(widget.Clickable),
	}

	a.ruleRadio.Value = "sum"

	return a
}

func (a *addRule) Layout(gtx Ctx, th *material.Theme) (nextScreen Screen) {
	nextScreen = a

	radio := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(0))
		in.Top = unit.Dp(64)
		return in.Layout(gtx, func(gtx Ctx) Dim {
			return layout.Flex{
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(material.RadioButton(th, a.ruleRadio, "sum", "Summa").Layout),
				layout.Rigid(material.RadioButton(th, a.ruleRadio, "set", "Par").Layout),
				layout.Rigid(material.RadioButton(th, a.ruleRadio, "single", "En tärning").Layout),
			)
		})
	}

	rolls := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(50))
		d2w := float32(0)
		if a.ruleRadio.Value == "set" {
			d2w = 1
		}
		return layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			FlexedInset(in, 1, func(gtx Ctx) Dim {
				if a.ruleRadio.Value == "sum" {
					return a.d1.Layout(gtx, 2, 12)
				}
				return a.d1.Layout(gtx, 1, 6)
			}),
			FlexedInset(in, d2w, func(gtx Ctx) Dim {
				if a.ruleRadio.Value == "set" {
					return a.d2.Layout(gtx, 1, 6)
				}
				return Dim{}
			}),
		)
	}

	text := func(gtx Ctx) Dim {
		edit := material.Editor(th, a.nameEdit, "Regel")
		edit.TextSize = material.H5(th, "").TextSize
		return edit.Layout(gtx)
	}

	save := func(gtx Ctx) Dim {
		saveBttn := material.Button(th, a.saveClick, "\nSpara\n")
		for a.saveClick.Clicked() {
			if newRule := a.saveRule(); newRule != nil {
				nextScreen = gameScreen(append(a.rules, newRule))
			}
		}
		return saveBttn.Layout(gtx)
	}

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx Ctx) Dim {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			layout.Rigid(radio),
			layout.Rigid(rolls),
			layout.Flexed(1, text),
			layout.Rigid(save),
		)
	})

	return nextScreen
}

func (a *addRule) saveRule() Rule {
	switch a.ruleRadio.Value {
	case "sum":
		return SumRule{
			Name: a.nameEdit.Text(),
			Sum:  a.d1.val,
		}
	case "set":
		return SetRule{
			Name: a.nameEdit.Text(),
			Set:  Roll{a.d1.val, a.d2.val},
		}
	case "single":
		return SingleRule{
			Name: a.nameEdit.Text(),
			Dice: a.d1.val,
		}
	}

	return nil
}

func (a *addRule) Rules() []Rule {
	return a.rules
}
