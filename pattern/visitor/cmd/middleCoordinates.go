package main

import (
	"L2/pattern/visitor/pkg"
	"fmt"
)

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) VisitForSquare(s *pkg.Square) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) VisitForCircle(c *pkg.Circle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) VisitForRectangle(r *pkg.Rectangle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for rectangle")
}
