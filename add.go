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
	saveButton = new(widget.Button)
)

func addRule(gtx *layout.Context, th *material.Theme) {
	radio := func() {
		in := layout.UniformInset(unit.Dp(0))
		in.Top = unit.Dp(64)
		in.Layout(gtx, func() {
			layout.Flex{
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(func() {
					th.RadioButton("sum", "Summa").Layout(gtx, ruleRadio)
				}),
				layout.Rigid(func() {
					th.RadioButton("set", "Par").Layout(gtx, ruleRadio)
				}),
				layout.Rigid(func() {
					th.RadioButton("single", "En tärning").Layout(gtx, ruleRadio)
				}),
			)
		})
	}

	rolls := func() {
		edit := th.Editor("0")
		edit.TextSize = fontSize
		in := layout.UniformInset(unit.Dp(50))
		layout.Flex{
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			layout.Rigid(func() {
				in.Layout(gtx, func() {
					edit.Layout(gtx, d1Edit)
				})
			}),
			layout.Rigid(func() {
				if ruleRadio.Value(gtx) == "set" {
					in.Layout(gtx, func() {
						edit.Layout(gtx, d2Edit)
					})
				}
			}),
		)
	}

	text := func() {
		edit := th.Editor("Regel")
		edit.TextSize = bigFont
		edit.Layout(gtx, nameEdit)
	}

	save := func() {
		bttn := th.Button("\nSpara\n")
		bttn.TextSize = fontSize
		for saveButton.Clicked(gtx) {
			if saveRule(gtx) == nil {
				playing = true
				d1Edit.SetText("")
				d2Edit.SetText("")
				nameEdit.SetText("")
			}
		}
		bttn.Layout(gtx, saveButton)
	}

	layout.UniformInset(unit.Dp(8)).Layout(gtx, func() {
		layout.Flex{
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

func saveRule(gtx *layout.Context) error {
	d1, err := strconv.Atoi(d1Edit.Text())
	if err != nil {
		return err
	}
	d2, err := strconv.Atoi(d2Edit.Text())
	if err != nil && ruleRadio.Value(gtx) == "set" {
		return err
	}

	var nrule Rule
	if ruleRadio.Value(gtx) == "sum" && d1 > 1 {
		nrule = SumRule{
			Name: nameEdit.Text(),
			Sum:  d1,
		}
	} else if ruleRadio.Value(gtx) == "set" && d1 > 0 && d1 < 7 && d2 > 0 && d2 < 7 {
		nrule = SetRule{
			Name: nameEdit.Text(),
			Set:  Roll{d1, d2},
		}
	} else if ruleRadio.Value(gtx) == "single" && d1 > 0 && d1 < 7 {
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
