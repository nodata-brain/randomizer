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

	s, err := random.Randomizer(*str, *n)
	if err != nil {
		fmt.Printf("error:%s", err)
		return
	}
	fmt.Println(s)
}
