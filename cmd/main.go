package main

import (
	//	"flag"
	"fmt"

	"randomizer/pkg/random"
)

func main() {
	s := random.Randomizer("int", nil, nil)
	fmt.Println(s)
}
