package main

import (
	"fmt"
)

func main() {
	// making zero value useful
	//1.
	var counter int
	fmt.Println(counter)
	for range "string" {
		counter++
	}
	println(counter)

	//2.
	var s []string
	s = append(s, "Hello")
	s = append(s, "World")
	fmt.Println(s)

	//3.
	var headers map[string]string
	fmt.Println(len(headers))

	for k, v := range headers { // iterates zero times
		_, _ = k, v
		panic("won't be executed")
	}

}
