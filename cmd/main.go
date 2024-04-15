package main

import (
	"fmt"
	"image"
	"os"
	sobelfilter "sobel/internal/sobel_filter"

	"image/color"
	"image/png"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	if len(os.Args) < 2 {
		fmt.Println("No image path")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	image.NewRGBA(bounds)

	new := image.NewRGBA(bounds)
	_ = new
	colors := make([][]uint8, bounds.Max.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		colors[y] = make([]uint8, bounds.Max.X)
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			colors[y][x] = uint8(int(float64(r)*0.3+float64(g)*0.59+float64(b)*0.11) >> 8)
		}
	}

	newColors := sobelfilter.Process(colors)
	fmt.Println(len(colors), len(colors[0]))
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			new.Set(x, y, color.RGBA{
				newColors[y][x],
				newColors[y][x],
				newColors[y][x],
				uint8(a >> 8),
			})
		}
	}

	out, err := os.Create(os.Args[1] + "_filter.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	png.Encode(out, new)

}
