package main

import (
	"fmt"
	"time"
)

/*
1. Find the issue with below code. Understand the root cause of it.
2. Fix the issue with using select `cancellation` or `timeout`.
*/
func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		select {
		case strings := <-strings:
			for s := range strings {
				// do something interesting
				fmt.Println(s)
			}
		case <-time.After(10 * time.Millisecond):
			fmt.Println("doWork exited.")
			close(completed)
		}

		return completed
	}

	// ch := make(chan string)
	// strings := "stringsdsdwse"
	// go func() {
	// 	ch <- strings
	// }()
	// doWork(ch)
	doWork(nil)
	fmt.Println("Done")
}
