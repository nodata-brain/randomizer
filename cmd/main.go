package main

import (
	//	"flag"
	"fmt"

	"randomizer/pkg/random"
)

func main() {
	s := random.Randomizer("int", 0)
	fmt.Println(s)
}
