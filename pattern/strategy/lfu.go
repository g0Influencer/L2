package strategy

import "fmt"

//конкретная стратегия
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}
