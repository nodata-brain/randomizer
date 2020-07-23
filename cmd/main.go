package main

import (
	"flag"
	"fmt"

	"randomizer/pkg/random"
)

func main() {
	str := flag.String("m", "", "mode randomizer /n int = 0~n /n str hoge")
	n := flag.Int("n", 0, "set range /n int n /n str word")
	flag.Parse()

	s := random.Randomizer(*str, *n)
	fmt.Println(s)
}
