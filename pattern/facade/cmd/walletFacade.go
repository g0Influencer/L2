package main

import (
	"L2/pattern/facade/pkg"
	"fmt"
)

// 	WalletFacade в данном примере является фасадом, предоставляющим простой интерфейс к системе классов оплаты.
//	Клиенту не нужно знать внутреннее устройство механизмов оплаты,ему достаточно просто отправить реквизиты карты,
//	CVV номер,стоимость оплаты и тип операции.

type WalletFacade struct {
	account *pkg.Account
	wallet *pkg.Wallet
	securityCode *pkg.SecurityCode
	ledger *pkg.Ledger
	notification *pkg.Notification
}

func newWalletFacade(accountID string, code int) *WalletFacade{
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		account: pkg.NewAccount(accountID),
		wallet: pkg.NewWallet(),
		securityCode: pkg.NewSecurityCode(code),
		ledger: &pkg.Ledger{},
		notification: &pkg.Notification{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting add money to wallet")
	err:= w.account.CheckAccount(accountID)
	if err!= nil{
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err!= nil{
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID,"credit",amount)
	return nil
}
func (w * WalletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error{
	fmt.Println("Starting debit money from wallet")
	err:= w.account.CheckAccount(accountID)
	if err!= nil{
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err!= nil{
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err!= nil{
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID,"debit",amount)
	return nil

}



