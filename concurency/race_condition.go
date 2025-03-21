package concurency

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    *sync.RWMutex
	count int
}

func (c *Counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count

}

func RaceCondition() {
	wg := &sync.WaitGroup{}
	c := Counter{
		mu: new(sync.RWMutex),
	}
	num := 100
	//var res int
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			c.inc(wg)

		}()
	}
	wg.Wait()

	fmt.Println("Counter: ", c.value())

}
