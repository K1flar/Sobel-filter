package sobelfilter

import "math"

var (
	kernelX = [][]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	kernelY = [][]int{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}
)

func Process(colors [][]uint8) [][]uint8 {
	h := len(colors)
	w := len(colors[0])
	res := make([][]uint8, h)

	for y := 1; y < h-1; y++ {
		res[y] = make([]uint8, w)
		for x := 1; x < w-1; x++ {
			gx, gy := 0, 0

			// Применяем ядра фильтра Собела к пикселям изображения
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					gx += int(colors[y+k][x+l]) * kernelX[k+1][l+1]
					gy += int(colors[y+k][x+l]) * kernelY[k+1][l+1]
				}
			}

			res[y][x] = uint8(math.Sqrt(float64(gx*gx + gy*gy)))
		}
	}

	return res
}
