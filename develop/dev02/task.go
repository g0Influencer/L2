package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	b, err := Unpack(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
// Unpack - распаковать строку
func Unpack(s string) (str string, err error) {
	// приводим к нижнему регистру
	s = strings.ToLower(s)
	// проверка на пустую строку
	if s == "" {
		return str, nil
	}
	// проверка на первый элемент, если он не в пределе алфавита
	if s[0] < 97 || s[0] > 122 {
		return str, fmt.Errorf("incorrent first element")
	}
	// цикл по символам строки
	for i := 0; i <= len(s)-1; i++ {
		// Если символ в пределах алфавита, то добавляем букву к итоговой строке
		if s[i] >= 97 && s[i] <= 122 {
			str += string(s[i])
			// Если у нас есть escape последовательность, добавляем в итоговую строку следующий символ и увеличиваем номер итерации на 1
		} else if s[i] == '\\' {
			str += string(s[i+1])
			i++
			// здесь получаем число, повторения символов. С учетом того, что один символ добавили
		} else {
			l := make([]byte, 0)
			for n := i; n <= len(s)-1; n++ {
				if s[n] >= 48 && s[n] <= 57 {
					l = append(l, s[n])
				} else {
					break
				}
			}
			a, err := strconv.Atoi(string(l))
			if err != nil {
				panic(err)
			}
			y := strings.Repeat(string(s[i-1]), a-1)
			str += y
		}

	}
	return
}
