package main

import "fmt"

func main() {
	var (
		a float64
		b float64
		c float64
	)

	fmt.Scan(&a, &b)

	c = (a + b) / 2
	fmt.Printf("MEDIA = %.5f\n", c)
}
