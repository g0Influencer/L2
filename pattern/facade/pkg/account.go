package pkg

import "fmt"

type Account struct {
	name string
}

func NewAccount(AccountName string) *Account{
	return &Account{
		name: AccountName,
	}
}

func (a *Account) CheckAccount(AccountName string) error{
	if a.name != AccountName{
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil

}