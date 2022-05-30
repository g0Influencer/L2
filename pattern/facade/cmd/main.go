package main

import (
	"fmt"
	"log"
)

/*	Фасад - структурный паттерн, который предоставляет простой(но урезанный) интерфейс к сложной системе объектов,
библиотеке или фреймворку.
	Фасад позволяет снизить общую сложность программы,помогает вынести код,зависимый от внешней системы в единственное место.

	Применимость:
		- Когда нам нужно представить простой или урезанный интерфейс к сложной подсистеме
		- Когда мы хотим разложить подсистему на отдельные слои
*/

func main(){
	walletFacade:= newWalletFacade("abc", 1234)
	fmt.Println()
	err:= walletFacade.addMoneyToWallet("abc",1234,100)
	fmt.Println()
	if err!= nil{
		log.Fatalf("Error: %develop\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc",1234,100)
	if err!= nil{
		log.Fatalf("Error: %develop\n", err.Error())
	}

}
