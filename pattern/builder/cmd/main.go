package main

import (
	"L2/pattern/builder/pkg"
	"fmt"
)

/*
	Паттерн Builder применяется,когда создаваемый объект большой и требует нескольких шагов.

	Строитель помогает вынести конструирование объекта за пределы его собственного класса,поручив это дело
	отдельным объектам,называемым строителями,тем самым мы избежим "телескопического конструктора".

	Строитель можно применить,если создание нескольких представлений объекта состоит из одинаковых этапов,
	которые отличаются в деталях.Интерфейс строителей определит все возможные этапы конструирования.Каждому представлению
	будет соответствовать собственный класс-родитель.А порядок этапов строительства будет задавать класс-директор.

 */

func main(){
	normalBuilder := pkg.GetBuilder("normal")
	iglooBuilder := pkg.GetBuilder("igloo")

	director := pkg.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %develop\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %develop\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %develop\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %develop\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
}
