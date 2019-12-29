package main

import (
	"fmt"
	"image/color"
)

type point struct {
	X, Y float64
}

func (p *point) setX(x float64) {
	p.X = x
}

func (p *point) setY(y float64) {
	p.Y = y
}

func (p point) print() {
	fmt.Printf("p.X = %f, p.Y = %f\n", p.X, p.Y)
}

type coloredPoint struct {
	point
	clr color.RGBA
}

func (clr coloredPoint) print() {
	fmt.Printf("I'm coloredPoint::print\n")
	fmt.Printf("clr.X = %f, clr.Y = %f\n", clr.X, clr.Y)
	clr.point.print()
}

func main() {
	fmt.Println("Method Inherit")
	var cp coloredPoint
	cp.setX(1.01)
	cp.setY(2.02)
	cp.print()
}
