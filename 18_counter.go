package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	num int
}

func (c *Counter) Inc(numGoroutin, finish int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		c.Lock()
		if c.num >= finish {
			c.Unlock()
			break
		}
		c.num++
		//fmt.Println("Inc by:", numGoroutin, "Counter:", c.num)
		c.Unlock()
		time.Sleep(1 * time.Microsecond)
	}
}

func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.num
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	finish := 1900
	numGoroutin := 4

	wg.Add(numGoroutin)
	for i := 0; i < numGoroutin; i++ {
		go func(i int) {
			counter.Inc(i, finish, &wg)
		}(i)
	}

	wg.Wait()

	fmt.Printf("Counter value: %d\n", counter.Value())
}
