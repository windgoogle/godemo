package main

import (
	"flag"
	"fmt"
)

var (
	x = flag.Int("x", 10, "input x")
	y = flag.Int("y", 20, "input y")
	s = flag.String("s", "hello world", "input string")
)

func main() {
	flag.Parse()

	fmt.Printf("x:%d, y:%d, s:%s", *x, *y, *s)
}
