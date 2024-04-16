package bwfilter

import (
	"image"
)

func Process(img image.Image, res *[][]uint8) {
	bounds := img.Bounds()

	w := bounds.Max.X
	h := bounds.Max.Y

	if (*res) == nil {
		(*res) = make([][]uint8, h)
	}

	for y := 0; y < h; y++ {
		if (*res)[y] == nil {
			(*res)[y] = make([]uint8, w)
		}
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			(*res)[y][x] = uint8(int(float64(r)*0.3+float64(g)*0.59+float64(b)*0.11) >> 8)
		}
	}
}
