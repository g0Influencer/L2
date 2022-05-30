package main

import "L2/pattern/command/pkg"

/*
	Комманда - поведенческий паттерн проектирования,позволяющий заворачивать запросы или простые операции в отдельные объекты.

	Применяется:
		- Когда нужно параметризировать объекты выполняемым действием
		- Когда нужно ставить операции в очередь,выполнять их по расписанию или передавать по сети
		- Когда нужна операция отмены
 */

func main(){
	tv:=&pkg.Tv{}
	onCommand:= &pkg.OnCommand{
		tv,
	}
	offCommand:= &pkg.OffCommand{
		tv,
	}
	onButton:=&pkg.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton:=&pkg.Button{
		Command: offCommand,
	}
	offButton.Press()

}
