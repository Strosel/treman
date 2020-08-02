package main

import (
	"gioui.org/layout"
)

//Ctx is a helper alias for less wide code
type Ctx = layout.Context

//Dim is a helper alias for less wide code
type Dim = layout.Dimensions

func RigidInset(in layout.Inset, widget layout.Widget) layout.FlexChild {
	return layout.Rigid(func(gtx Ctx) Dim {
		return in.Layout(gtx, widget)
	})
}
