package main

import (
	"image/color"

	"golang.org/x/image/colornames"
)

var (
	BLACK          = color.NRGBAModel.Convert(colornames.Black).(color.NRGBA)
	WHITE          = color.NRGBAModel.Convert(colornames.White).(color.NRGBA)
	MEDIUMSEAGREEN = color.NRGBAModel.Convert(colornames.Mediumseagreen).(color.NRGBA)
	ROSYBROWN      = color.NRGBAModel.Convert(colornames.Rosybrown).(color.NRGBA)
	ROYALBLUE      = color.NRGBAModel.Convert(colornames.Royalblue).(color.NRGBA)
	MYRED          = color.NRGBAModel.Convert(color.RGBA{230, 0, 0, 255}).(color.NRGBA)
)
