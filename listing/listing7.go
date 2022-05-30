package main

import (
	"fmt"
	"math/rand"
	"time"
)
// При запуске программы мы получим сначала данные,которые ожидаем получить,а затем попадем в бесконечный цикл,в котором
// будут выводиться 0. Это происходит из-за того,что мы в бесконечном цикле пытаемся получить данные из канала даже если
// он закрыт,а значением по умолчанию для закрытого целочисленного канала является 0. Чтобы избежать данной ошибки, нам
// нужно проверить в каждом case закрыт ли канал,из которого мы получаем значение.


func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)


	go func(c chan int) {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}(c)
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}

