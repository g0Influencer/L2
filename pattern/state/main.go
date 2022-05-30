package state

import (
	"fmt"
	"log"
)


/*
		Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение
		в зависимости от своего состояния.Извне создаётся впечатление, что изменился класс объекта

		Применимость:
			- Когда у нас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
			  причём типов состояний много, и их код часто меняется

			- Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
			  которые выбирают поведения в зависимости от текущих значений полей класса

			- Когда мы сознательно используете табличную машину состояний, построенную на условных операторах,
			  но вынуждены мириться с дублированием кода для похожих состояний и переходов
 */
func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
