package pkg

import "fmt"

type Ledger struct {

}

func (s *Ledger) MakeEntry(accountID, txnType string, amount int){
	fmt.Printf("Make Ledger entry for accountId %develop with txnType %develop for amount %d", accountID, txnType, amount)
	return
}
