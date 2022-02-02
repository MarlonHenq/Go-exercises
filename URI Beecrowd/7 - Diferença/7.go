package main

import "fmt"

func main() {
	var (
		a float64
		b float64
		c float64
		d float64
		e float64
	)

	fmt.Scan(&a, &b, &c, &d)

	a = a * b
	c = c * d
	e = a - c

	fmt.Printf("DIFERENCA = %.0f\n", e)
}
