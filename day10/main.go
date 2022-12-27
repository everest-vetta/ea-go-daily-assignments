package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
)

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	fmt.Println("Hello World")
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go counter(&wg)
	}
	wg.Wait()
}

func counter(wg *sync.WaitGroup) {
	println("1")
	wg.Done()
}
