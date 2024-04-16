package app

import (
	"fmt"
	"image"
	"image/color"
	bwfilter "sobel/internal/filters/bw_filter"
	sobelfilter "sobel/internal/filters/sobel_filter"
	"sync"
	"time"
)

func Run(img image.Image, numOfThreads int) (*image.RGBA, error) {
	bounds := img.Bounds()
	w := bounds.Max.X
	h := bounds.Max.Y
	countStr := h / numOfThreads

	gray := make([][]uint8, h)
	bwfilter.Process(img, &gray)

	gradient := make([][]uint8, h)
	startY := 1
	endY := countStr - 1
	wg := sync.WaitGroup{}
	s := time.Now()
	for t := 0; t < numOfThreads; t++ {
		wg.Add(1)
		go func(startY, endY, t int) {
			defer wg.Done()
			sobelfilter.Process(startY, endY, gray, &gradient)
		}(startY, endY, t)

		startY = endY
		if t+1 == numOfThreads-1 {
			endY = h - 1
		} else {
			endY = startY + countStr
		}
	}
	wg.Wait()
	fmt.Println(time.Since(s))

	new := image.NewRGBA(bounds)
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			new.Set(x, y, color.RGBA{
				gradient[y][x],
				gradient[y][x],
				gradient[y][x],
				uint8(a >> 8),
			})
		}
	}

	return new, nil
}
