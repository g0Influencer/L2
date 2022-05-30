package factory_method

import "fmt"

/*Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс
  для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

   Применимость:
    - Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать наш код.
 	- Когда мы хотим дать возможность пользователям расширять части вашего фреймворка или библиотеки
    - Когда мы хотим экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых

*/
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}