package main

import (
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/markbates/pkger"
)

func loadFont() ([]text.FontFace, error) {
	f, err := pkger.Open("/assets/dice.ttf")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	diceFont := make([]byte, 6000)
	_, err = f.Read(diceFont)
	if err != nil {
		return nil, err
	}

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
