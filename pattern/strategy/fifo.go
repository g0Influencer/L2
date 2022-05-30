package strategy

import "fmt"


// конкретная стратегия
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}
