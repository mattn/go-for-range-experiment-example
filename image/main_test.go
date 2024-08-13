package main

import (
	"fmt"
	"image"
)

func ExamplePoints() {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 2, Y: 5},
		Max: image.Point{X: 5, Y: 8},
	})

	for p := range Points(img) {
		fmt.Println(p)
		// eg. r, g, b, a := img.At(p.X, p.Y).RGBA()
	}

	// Output: (2,5)
	// (3,5)
	// (4,5)
	// (2,6)
	// (3,6)
	// (4,6)
	// (2,7)
	// (3,7)
	// (4,7)
}
