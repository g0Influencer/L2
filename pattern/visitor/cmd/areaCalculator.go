package main

import (
	"L2/pattern/visitor/pkg"
	"fmt"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) VisitForSquare(s *pkg.Square){
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}
func (a *areaCalculator) VisitForCircle(c *pkg.Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) VisitForRectangle(r *pkg.Rectangle) {
	fmt.Println("Calculating area for rectangle")
}


