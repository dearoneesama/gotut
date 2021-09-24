package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 20, 20)
}
func (im Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
