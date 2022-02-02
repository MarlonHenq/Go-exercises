package main

import "fmt"

func main() {
	var (
		a float64
		b float64
		c float64
		d float64
	)

	fmt.Scan(&a, &b, &c)

	a = a * 2
	b = b * 3
	c = c * 5

	d = (a + b + c) / 10
	fmt.Printf("MEDIA = %.1f\n", d)
}
