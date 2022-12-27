package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count int64
}

func (numCounter *counter) Add(inpNumber int) {
	_ = atomic.AddInt64(&numCounter.count, int64(inpNumber))
}

// A simple program to trigger 10 goroutines
// & each to add 100 to Counter via iteration one by one
func main() {
	var numCounter counter

	group := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				numCounter.Add(1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(numCounter.count)
}
