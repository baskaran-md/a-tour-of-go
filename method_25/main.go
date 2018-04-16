package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	W int
	H int
}

func (img Image) ColorModel() color.Model {
	return color.NRGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.W, img.H)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y + y ^ x)
	return color.RGBA{
		R: v,
		G: v,
		B: 255,
		A: 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
