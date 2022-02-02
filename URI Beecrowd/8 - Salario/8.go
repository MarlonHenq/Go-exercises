package main

import "fmt"

func main() {
	var (
		number  int
		salario float64
		horas   float64
	)

	fmt.Scan(&number)
	fmt.Scan(&horas, &salario)

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", salario*horas)
}
