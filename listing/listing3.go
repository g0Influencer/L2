package main

import (
	"fmt"
	"os"
)
//  error представляет собой интерфейс. Foo возвращает указатель на структуру PathError, которая реализует интерфейс error.
//	Получается, что интерфейс не пустой, в нем лежит структура, которая уже содержит nil
//	Поэтому выводится значение nil, но интерфейс не является nil

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err) // nil
	fmt.Println(err == nil) // false
}

