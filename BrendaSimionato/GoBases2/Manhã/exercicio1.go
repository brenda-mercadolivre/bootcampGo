package main

import "fmt"

func calcImposto(salario float64) float64 {
	if salario <= 50000.00 {
		return 0.0
	} else if salario <= 150000.00 {
		return salario * 0.17
	}
	return salario * 0.27
}

func main() {
	fmt.Printf("Imposto de até R$50.000: %.2f\n",
		calcImposto(50000))
	fmt.Printf("Imposto de até R$150.000: %.2f\n",
		calcImposto(150000))
	fmt.Printf("Imposto mais que R$150.000: %.2f\n",
		calcImposto(150001))
}
