package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitSquare(*s)
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitCircle(*c)
}

type triangle struct {
	sides [3]int
}

func (t *triangle) accept(v visitor) {
	v.visitTriangle(*t)
}

type visitor interface {
	visitSquare(square)
	visitCircle(circle)
	visitTriangle(triangle)
}

type areaCalculator struct {
	area float64
}

func (ac *areaCalculator) visitSquare(s square) {
	ac.area = float64(s.side * s.side)
}

func (ac *areaCalculator) visitCircle(c circle) {
	ac.area = math.Pi * float64(c.radius*c.radius)
}

func (ac *areaCalculator) visitTriangle(t triangle) {
	var sp float64 = float64(t.sides[0]+t.sides[1]+t.sides[2]) / 2
	// heron's formula
	ac.area = math.Sqrt(sp * (sp - float64(t.sides[0])) * (sp - float64(t.sides[1])) * (sp - float64(t.sides[2])))
}

func (ac *areaCalculator) calculate() float64 {
	return ac.area
}

func VisitorExample() {
	s := square{side: 3}
	c := circle{radius: 2}
	t := triangle{sides: [...]int{3, 4, 5}}

	ac := areaCalculator{}

	s.accept(&ac)
	fmt.Println("Square area:", ac.calculate())

	c.accept(&ac)
	fmt.Println("Circle area:", ac.calculate())

	t.accept(&ac)
	fmt.Println("Triangle area:", ac.calculate())
}
