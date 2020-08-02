package main

import (
	"gioui.org/layout"
)

func RigidInset(in layout.Inset, widget layout.Widget) layout.FlexChild {
	return layout.Rigid(func(gtx Ctx) Dim {
		return in.Layout(gtx, widget)
	})
}
