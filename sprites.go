package main

import (
	"image"
	"image/draw"
	"image/png"

	"gioui.org/op/paint"
	"github.com/markbates/pkger"
)

func subImage(img image.Image, rect image.Rectangle) image.Image {
	out := image.NewRGBA(image.Rectangle{
		Max: rect.Size(),
	})

	draw.Draw(out, out.Bounds(), img, rect.Min, draw.Src)

	return out
}

func loadSprites() ([]paint.ImageOp, error) {
	f, err := pkger.Open("/sprites.png")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pic, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	sprites := []paint.ImageOp{}

	i := 0
	b := pic.Bounds()
	dy := b.Dy() / 5
	dx := b.Dx() / 3
	for y := 0; y < b.Dy(); y += dy {
		for x := 0; x < b.Dx(); x += dx {
			if i < 15 {
				img := subImage(pic, image.Rect(x, y, x+dx, y+dy))
				sprites = append(sprites, paint.NewImageOp(img))
				i++
			}
		}
	}

	return sprites, nil
}
