package main

import "fmt"

func main() {
	const PI float64 = 3.14159
	var (
		area float64
		raio float64
	)

	fmt.Scanln(&raio)
	area = PI * raio * raio

	fmt.Printf("A=%.4f\n", area)
}
