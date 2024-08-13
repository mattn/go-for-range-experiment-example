package main

import (
	"image"
	"iter"
)

func Points(img image.Image) iter.Seq[image.Point] {
	return func(yield func(image.Point) bool) {
		bounds := img.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if !yield(image.Point{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
