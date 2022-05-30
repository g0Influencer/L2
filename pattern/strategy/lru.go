package strategy

import "fmt"

//конкретная стратегия
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}
