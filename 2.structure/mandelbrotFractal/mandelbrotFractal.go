package mandelbrotFractal

import (
	"image"
	"image/color"
	"math/cmplx"
)

func PrintFractal() *image.RGBA {
	const (
		xMax, xMin, yMax, yMin = +2, -2, +2, -2
		width, height          = 4096, 4096
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}
func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		constrast  = 15
	)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - n*constrast}
		}
	}
	return color.Black
}
