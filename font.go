package main

import (
	_ "embed"

	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
)

//go:embed assets/dice.ttf
var diceFont []byte

func loadFont() ([]text.FontFace, error) {
	face, err := opentype.Parse(diceFont)
	if err != nil {
		return nil, err
	}

	return append(gofont.Collection(), text.FontFace{
		Font: text.Font{
			Typeface: "Go",
			Variant:  "Dice",
		},
		Face: face}), nil
}
