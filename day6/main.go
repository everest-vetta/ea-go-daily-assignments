package main

import (
	"fmt"
	"time"
)

func Sample(routineID int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("RountineID: %v i: %v\n", routineID, i)
		time.Sleep(time.Second * 30)
	}
}

func main() {
	//main()
	go Sample(1)
	go Sample(2)
	go Sample(3)
}
