package main

import "L2/pattern/chainOfResp/pkg"

/*
	Цепочка обязанностей - это поведенческий паттерн проектирования,который позволяет передавать запросыпоследовательно
	по цепочке обработчиков.Каждый последующий обработчик решает,может ли он обработать запрос сам
	и стоит ли передавать запрос дальше по цепи.

	Применяется:
		- Когда  программа должна обрабатывать разнообразные запросы несколькими способами,
		  но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
		- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
		- Когда набор объектов, способных обработать запрос, должен задаваться динамически.

 */

func main(){
	cashier := &pkg.Cashier{}

	//Set next for medical department
	medical := &pkg.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &pkg.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &pkg.Reception{}
	reception.SetNext(doctor)

	patient := &pkg.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
