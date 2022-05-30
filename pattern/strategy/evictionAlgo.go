package strategy

// интерфейс стратегии
type evictionAlgo interface {
	evict(c *cache)
}