package main

import (
	_ "embed"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
)

//go:embed assets/dice.ttf
var rawDiceFont []byte
var diceFont font.Font

func loadFont() ([]text.FontFace, error) {
	face, err := opentype.Parse(rawDiceFont)
	if err != nil {
		return nil, err
	}

	diceFont = font.Font{
		Typeface: "Go",
		Variant:  "Dice",
	}

	return append(gofont.Collection(), text.FontFace{
		Font: diceFont,
		Face: face}), nil
}
