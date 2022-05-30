package main

import (
	"L2/pattern/visitor/pkg"
	"fmt"
)

/*
	Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
не изменяя типы этих объектов, над которыми эти операции могут выполняться.

Применяется:
	Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов
	Когда новое поведение нужно только для некоторых классов из существующих

Плюсы и минусы:
	+ Упрощает добавление операций, работающих со сложными структурами объектов.
	+ Посетитель может накапливать состояние при обходе структуры элементов

	- Паттерн не оправдан, если иерархия элементов часто меняется
 */

func main(){
	square := &pkg.Square{Side: 2}
	circle := &pkg.Circle{Radius: 3}
	rectangle := &pkg.Rectangle{L: 2, B: 3}

	areaCalculator := &areaCalculator{}
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)


}
