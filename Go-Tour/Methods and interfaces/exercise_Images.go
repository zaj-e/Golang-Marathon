package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{
	cm color.Model
	r image.Rectangle
	cl color.Color
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	i.r = image.Rect(0, 0, 100, 100)
	return i.r
}

func (i Image) At(x, y int) color.Color {
	n_x := uint8(x)
	n_y := uint8(y)
	v := n_x*n_y
	return color.RGBA{v,v,255,255}
}


func main() {
	m := Image{}
	pic.ShowImage(m)
}