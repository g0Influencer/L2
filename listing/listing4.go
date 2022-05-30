package main


// После вывода чисел от 0 до 9,которые мы передали в канал,мы получим deadlock,
//	потому что цикл for n := range ch получает значения из канала до тех пор, пока он не закрыт.
//

func main() {
	ch := make(chan int)
	go func() {
		 //defer close(ch)  - теперь дедлока не будет
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
