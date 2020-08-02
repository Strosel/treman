package main

import (
	"fmt"
	"strconv"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type addRule struct {
	rules []Rule

	ruleRadio *widget.Enum
	d1Edit    *widget.Editor
	d2Edit    *widget.Editor
	nameEdit  *widget.Editor
	saveClick *widget.Clickable
}

func addRuleScreen(rules []Rule) Screen {
	a := &addRule{
		rules:     rules,
		ruleRadio: new(widget.Enum),
		d1Edit:    new(widget.Editor),
		d2Edit:    new(widget.Editor),
		nameEdit:  new(widget.Editor),
		saveClick: new(widget.Clickable),
	}

	a.d1Edit.SingleLine = true
	a.d2Edit.SingleLine = true
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
		return layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			RigidInset(in, material.Editor(th, a.d1Edit, "0").Layout),
			RigidInset(in, func(gtx Ctx) Dim {
				if a.ruleRadio.Value == "set" {
					return material.Editor(th, a.d2Edit, "0").Layout(gtx)
				}
				return Dim{}
			}),
		)
	}

	text := func(gtx Ctx) Dim {
		edit := material.Editor(th, a.nameEdit, "Regel")
		edit.TextSize = bigFont
		return edit.Layout(gtx)
	}

	save := func(gtx Ctx) Dim {
		saveBttn := material.Button(th, a.saveClick, "\nSpara\n")
		for a.saveClick.Clicked() {
			if a.saveRule() == nil {
				nextScreen = gameScreen(a.rules)
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

func (a *addRule) saveRule() error {
	d1, err := strconv.Atoi(a.d1Edit.Text())
	if err != nil {
		return err
	}
	d2, err := strconv.Atoi(a.d2Edit.Text())
	if err != nil && a.ruleRadio.Value == "set" {
		return err
	}

	var nrule Rule
	if a.ruleRadio.Value == "sum" && d1 > 1 && d1 < 13 {
		nrule = SumRule{
			Name: a.nameEdit.Text(),
			Sum:  d1,
		}
	} else if a.ruleRadio.Value == "set" && d1 > 0 && d1 < 7 && d2 > 0 && d2 < 7 {
		nrule = SetRule{
			Name: a.nameEdit.Text(),
			Set:  Roll{d1, d2},
		}
	} else if a.ruleRadio.Value == "single" && d1 > 0 && d1 < 7 {
		nrule = SingleRule{
			Name: a.nameEdit.Text(),
			Dice: d1,
		}
	} else {
		return fmt.Errorf("Something goofed")
	}

	a.rules = append(a.rules, nrule)
	return nil
}

func (a *addRule) Rules() []Rule {
	return a.rules
}
