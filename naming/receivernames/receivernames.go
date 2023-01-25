package receivernames

import "math"

type Point struct {
	x, y float64
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
