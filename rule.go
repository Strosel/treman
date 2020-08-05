package main

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

//Rule is a rule
type Rule interface {
	String() string
	Valid(Roll) bool
	Widget(th *material.Theme) func(gtx Ctx) Dim
}

//SumRule is a rule based on the sum of a roll
type SumRule struct {
	Name, Desc string
	Sum        int
}

func (sr SumRule) String() string {
	return sr.Name
}

func (sr SumRule) Valid(r Roll) bool {
	return sr.Sum == (r[0] + r[1])
}

func (sr SumRule) Widget(th *material.Theme) func(gtx Ctx) Dim {
	return func(gtx Ctx) Dim {
		dice := material.Body1(th, fmt.Sprint(0))
		dice.Font.Variant = "Dice"

		return layout.Flex{}.Layout(gtx,
			layout.Rigid(dice.Layout),
			layout.Rigid(material.Body1(th, " + ").Layout),
			layout.Rigid(dice.Layout),
			layout.Rigid(material.Body1(th, fmt.Sprintf(" = %v\t%v\n%v", sr.Sum, sr.Name, sr.Desc)).Layout),
		)
	}
}

//SetRule is a rule based on a specific roll
type SetRule struct {
	Name, Desc string
	Set        Roll
}

func (sr SetRule) String() string {
	return sr.Name
}

func (sr SetRule) Valid(r Roll) bool {
	lr := (sr.Set[0] == r[0] && sr.Set[1] == r[1])
	rl := (sr.Set[1] == r[0] && sr.Set[0] == r[1])
	return lr || rl
}

func (sr SetRule) Widget(th *material.Theme) func(gtx Ctx) Dim {
	return func(gtx Ctx) Dim {
		dice := material.Body1(th, fmt.Sprint(sr.Set[0], sr.Set[1]))
		dice.Font.Variant = "Dice"

		return layout.Flex{}.Layout(gtx,
			layout.Rigid(dice.Layout),
			layout.Rigid(material.Body1(th, fmt.Sprintf("\t%v\n%v", sr.Name, sr.Desc)).Layout),
		)
	}
}

//SingleRule is a rule based on a single dice
type SingleRule struct {
	Name, Desc string
	Dice       int
}

func (sr SingleRule) String() string {
	return sr.Name
}

func (sr SingleRule) Valid(r Roll) bool {
	return sr.Dice == r[0] || sr.Dice == r[1]
}

func (sr SingleRule) Widget(th *material.Theme) func(gtx Ctx) Dim {
	return func(gtx Ctx) Dim {
		dice := material.Body1(th, fmt.Sprint(sr.Dice))
		dice.Font.Variant = "Dice"

		return layout.Flex{}.Layout(gtx,
			layout.Rigid(dice.Layout),
			layout.Rigid(material.Body1(th, fmt.Sprintf("\t%v\n%v", sr.Name, sr.Desc)).Layout),
		)
	}
}

type specialRule struct {
	Name, Desc string
	Rule       func(Roll) bool
	Wid        func(sr specialRule, th *material.Theme) func(gtx Ctx) Dim
}

func (sr specialRule) String() string {
	return sr.Name
}

func (sr specialRule) Valid(r Roll) bool {
	return sr.Rule(r)
}

func (sr specialRule) Widget(th *material.Theme) func(gtx Ctx) Dim {
	return sr.Wid(sr, th)
}

//default rules
func drules() []Rule {
	return []Rule{
		specialRule{
			Name: "Treman",
			Desc: "Treman dricker",
			Rule: func(r Roll) bool {
				//Fast än treman dricker på 3,3 (då det är ny treman) ska det inte stå "treman dricker och ny treman"
				return (r[0] == 3 || r[1] == 3) && (r[0] != r[1])
			},
			Wid: func(sr specialRule, th *material.Theme) func(gtx Ctx) Dim {
				return func(gtx Ctx) Dim {
					dice := material.Body1(th, "3")
					dice.Font.Variant = "Dice"

					return layout.Flex{}.Layout(gtx,
						layout.Rigid(dice.Layout),
						layout.Rigid(material.Body1(th, fmt.Sprintf("\t%v\n%v", sr.Name, sr.Desc)).Layout),
					)
				}
			},
		},
		SetRule{
			Name: "Krig",
			Desc: "Välj en annan spelare. Ni är nu i krig, dricker den ena så dricker bägge",
			Set:  Roll{1, 1},
		},
		SetRule{
			Name: "Utmaning",
			Desc: "Välj en annan spelare och vars en tärning. Den som slår högst blir ny treman.",
			Set:  Roll{1, 2},
		},
		SetRule{
			Name: "En ferrari",
			Desc: "Sist att låtsas köra bil dricker. (\"Dark humour\" variant finns)",
			Set:  Roll{1, 4},
		},
		SetRule{
			Name: "Ny Treman",
			Desc: "Grattis! Du är nu treman",
			Set:  Roll{3, 3},
		},
		SetRule{
			Name: "Jag har aldrig sett...",
			Desc: "Häfv resten av din enhet och skapa en ny regel eller dela ut 6+6 klunkar.",
			Set:  Roll{6, 6},
		},
		SetRule{
			Name: "Dela ut 2+2 klunkar",
			Set:  Roll{2, 2},
		},
		SetRule{
			Name: "Dela ut 4+4 klunkar",
			Set:  Roll{4, 4},
		},
		SetRule{
			Name: "Dela ut 5+5 klunkar",
			Set:  Roll{5, 5},
		},
		SumRule{
			Name: "Seven ahead",
			Desc: "Personen framför den som slår i ordningen dricker.",
			Sum:  7,
		},
		SumRule{
			Name: "Nine behind",
			Desc: "Personen bakom den som slår i ordningen dricker.",
			Sum:  9,
		},
		SumRule{
			Name: "Finger på näsan",
			Desc: "Sist att sätta fingret på näsan dricker.",
			Sum:  11,
		},
	}
}
