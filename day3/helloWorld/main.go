package main

import (
	"fmt"

	"rsc.io/sampler"
)

func main() {

	fmt.Print("Calling sampler.Hello() :: ")
	fmt.Println(sampler.Hello())
	fmt.Print("Calling sampler.Glass() :: ")
	fmt.Println(sampler.Glass())
	// fmt.Println(sampler.Go())
	// fmt.Println(sampler.Opt())

}
