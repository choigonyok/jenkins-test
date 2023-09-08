package main

import (
	"fmt"
	"test/sum"
)

func main() {
	a := 10
	b := 20
	c := sum.Sum(a, b)
	d := sum.Multi(a, b)

	fmt.Println(c, d)
}
