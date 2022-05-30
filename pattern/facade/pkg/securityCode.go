package pkg

import "fmt"

type SecurityCode struct {
	code int
}

func NewSecurityCode(newCode int) *SecurityCode{
	return &SecurityCode{
		code: newCode,
	}
}

func (sc *SecurityCode) CheckCode(code int) error {
	if code != sc.code{
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
