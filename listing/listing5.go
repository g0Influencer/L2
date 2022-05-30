package main


// Программа выведет "error". Функция test() возвращает указатель на экземпляр структуры customError,
// а в функции main() мы сравниваем тип *customError с nil,что некорректно.

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func Test() *customError { // если поменять тип возвращаемого значения на error,то получим ожидаемый результат
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = Test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

