package dev01

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func CurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Request time error : %s", err.Error())
	}
	fmt.Printf("Сurrent time: %s", time)
}
