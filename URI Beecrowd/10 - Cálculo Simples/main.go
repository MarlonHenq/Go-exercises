package main

import "fmt"

func main() {
	var (
		codigo1, codigo2 int
		num1, num2       int
		price1, price2   float32
	)

	fmt.Scan(&codigo1)
	fmt.Scan(&num1)
	fmt.Scan(&price1)

	fmt.Scan(&codigo2)
	fmt.Scan(&num2)
	fmt.Scan(&price2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", (float32(num1)*price1)+(float32(num2)*price2))
}
