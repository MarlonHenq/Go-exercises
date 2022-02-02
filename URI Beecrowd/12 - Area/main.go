package main

import "fmt"

func main() {
	var (
		a, b, c float64
	)

	const pi = 3.14159

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Printf("TRIANGULO: %.3f\n", a*c/2)
	fmt.Printf("CIRCULO: %.3f\n", pi*c*c)
	fmt.Printf("TRAPEZIO: %.3f\n", ((a+b)/2)*c)
	fmt.Printf("QUADRADO: %.3f\n", b*b)
	fmt.Printf("RETANGULO: %.3f\n", a*b)
}
