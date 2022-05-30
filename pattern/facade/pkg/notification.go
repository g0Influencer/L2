package pkg

import "fmt"

type Notification struct {

}

func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit Notification")
	return
}

func (n *Notification) SendWalletDebitNotification(){
	fmt.Println("Sending wallet debit Notification")
	return
}