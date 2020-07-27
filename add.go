package main

import (
	"fmt"
	"strconv"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	ruleRadio  = new(widget.Enum)
	d1Edit     = new(widget.Editor)
	d2Edit     = new(widget.Editor)
	nameEdit   = new(widget.Editor)
	saveButton = new(widget.Clickable)
)

func addRule(gtx Ctx, th *material.Theme) {
	radio := func(gtx Ctx) Dim {
		in := layout.UniformInset(unit.Dp(0))
		in.Top = unit.Dp(64)
		return in.Layout(gtx, func(gtx Ctx) Dim {
			return layout.Flex{
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(material.RadioButton(th, ruleRadio, "sum", "Summa").Layout),
				layout.Rigid(material.RadioButton(th, ruleRadio, "set", "Par").Layout),
				layout.Rigid(material.RadioButton(th, ruleRadio, "single", "En tärning").Layout),
			)
		})
	}

	rolls := func(gtx Ctx) Dim {
		// edit.TextSize = fontSize
		in := layout.UniformInset(unit.Dp(50))
		return layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			layout.Rigid(func(gtx Ctx) Dim {
				return in.Layout(gtx, material.Editor(th, d1Edit, "0").Layout)
			}),
			layout.Rigid(func(gtx Ctx) Dim {
				if ruleRadio.Value == "set" {
					return in.Layout(gtx, material.Editor(th, d2Edit, "0").Layout)
				}
				return Dim{}
			}),
		)
	}

	text := func(gtx Ctx) Dim {
		edit := material.Editor(th, nameEdit, "Regel")
		edit.TextSize = bigFont
		return edit.Layout(gtx)
	}

	save := func(gtx Ctx) Dim {
		bttn := material.Button(th, saveButton, "\nSpara\n")
		bttn.TextSize = fontSize
		for saveButton.Clicked() {
			if saveRule() == nil {
				playing = true
				d1Edit.SetText("")
				d2Edit.SetText("")
				nameEdit.SetText("")
			}
		}
		return bttn.Layout(gtx)
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
}

func saveRule() error {
	d1, err := strconv.Atoi(d1Edit.Text())
	if err != nil {
		return err
	}
	d2, err := strconv.Atoi(d2Edit.Text())
	if err != nil && ruleRadio.Value == "set" {
		return err
	}

	var nrule Rule
	if ruleRadio.Value == "sum" && d1 > 1 {
		nrule = SumRule{
			Name: nameEdit.Text(),
			Sum:  d1,
		}
	} else if ruleRadio.Value == "set" && d1 > 0 && d1 < 7 && d2 > 0 && d2 < 7 {
		nrule = SetRule{
			Name: nameEdit.Text(),
			Set:  Roll{d1, d2},
		}
	} else if ruleRadio.Value == "single" && d1 > 0 && d1 < 7 {
		nrule = SingleRule{
			Name: nameEdit.Text(),
			Dice: d1,
		}
	} else {
		return fmt.Errorf("Something goofed")
	}
	if nrule != nil {
		rules = append(rules, nrule)
	}
	return nil
}
