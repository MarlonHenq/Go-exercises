package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
		c int
	)

	fmt.Scan(&a, &b)

	c = a + b
	fmt.Println("X =", c)

}
