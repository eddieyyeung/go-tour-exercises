package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width  int
	height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{123 + uint8(x), 123 + uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
