package main

import "gioui.org/widget/material"

type Screen interface {
	Layout(gtx Ctx, th *material.Theme) (nextScreen Screen)
	Rules() []Rule
}
