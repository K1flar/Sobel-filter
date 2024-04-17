package main

import (
	"fmt"
	"image"
	"os"
	"sobel/cmd/app"

	"image/png"
)

func main() {
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

	newImg, err := app.Run(img, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := os.Create(os.Args[1] + "_filter.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	png.Encode(out, newImg)

}
