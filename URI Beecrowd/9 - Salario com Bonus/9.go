package main

import "fmt"

func main() {
	var (
		nome     string
		salarioF float64
		vendas   float64
		bonus    float64
	)

	fmt.Scan(&nome)
	fmt.Scan(&salarioF)
	fmt.Scan(&vendas)

	bonus = vendas * 0.15

	salarioF += bonus

	fmt.Printf("TOTAL = R$ %.2f\n", salarioF)

}
