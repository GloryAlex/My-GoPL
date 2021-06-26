package surface

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6 //30º
)

var sinAngle, cosAngle = math.Sin(angle), math.Cos(angle)

func PrintSurface() string {
	var ret = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			ret += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	ret += "</svg>"
	return ret
}

func corner(i, j int) (float64, float64) {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	z := f(x, y)
	//if z > height {
	//	return 0, 0
	//}

	sx := width/2 + (x-y)*cosAngle*xyScale
	sy := height/2 + (x+y)*sinAngle*xyScale - z*zScale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
