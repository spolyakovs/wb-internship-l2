package main

import "fmt"

type Strategy interface {
	get(*Cache) int
}

type Cache struct {
	stack    []int
	strategy Strategy
}

func NewCache(s Strategy) *Cache {
	return &Cache{
		stack:    []int{1, 2, 3, 4, 5}, // fill it by default
		strategy: s,
	}
}

func (c *Cache) Put(elem int) {
	c.stack = append(c.stack, elem)
}

func (c *Cache) Pop() int {
	return c.strategy.get(c)
}

func (c *Cache) GetLen() int {
	return len(c.stack)
}

type Fifo struct{}

func (fifo Fifo) get(c *Cache) int {
	result := c.stack[0]
	c.stack = c.stack[1:]
	return result
}

type Lifo struct{}

func (lifo Lifo) get(c *Cache) int {
	result := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	return result
}

func StrategyExample() {
	fifoCache := NewCache(Fifo{})
	lifoCache := NewCache(Lifo{})

	fmt.Println("FIFO")
	for fifoCache.GetLen() > 0 {
		fmt.Printf("%v ", fifoCache.Pop())
	}

	fmt.Println("\nLIFO")
	for lifoCache.GetLen() > 0 {
		fmt.Printf("%v ", lifoCache.Pop())
	}
	fmt.Println()
}
