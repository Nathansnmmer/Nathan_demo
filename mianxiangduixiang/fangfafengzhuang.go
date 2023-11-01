package main

import (
	"fmt"
	"math"
)

func getDis(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}
func (this Point) getDis2(p Point) float64 {
	return math.Sqrt((this.x-p.x)*(this.x-p.x) + (this.y-p.y)*(this.y-p.y))
}

type Point struct {
	x, y float64
}

func main() {
	p1 := Point{0.0, 0.0}
	p2 := Point{3.0, 4.0}
	fmt.Println(getDis(p1, p2))
	fmt.Println(p2.getDis2(p1))
}
