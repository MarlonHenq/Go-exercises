package main

import "fmt"

func main() {
	var (
		raio, result float64
	)

	const pi = 3.14159

	fmt.Scan(&raio)

	raioCubo := raio * raio * raio

	result = (4.0 / 3) * pi * raioCubo

	fmt.Printf("VOLUME = %.3f\n", result)
}
