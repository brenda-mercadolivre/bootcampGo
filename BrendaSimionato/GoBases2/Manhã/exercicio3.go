package main

import (
	"errors"
	"fmt"
)

func calcSalario(categoria string, horas int) (float64, error) {
	if categoria == "C" {
		return float64(horas) * 1000, nil
	}
	if categoria == "B" {
		salario := float64(horas) * 1500
		if horas > 160 {
			return salario * 1.2, nil
		}
		return salario, nil
	}
	if categoria == "A" {
		salario := float64(horas) * 3000
		if horas > 160 {
			return salario * 1.5, nil
		}
		return salario, nil
	}
	return 0.0, errors.New("Categoria Inválida")
}
func main() {
	salario, _ := calcSalario("C", 162)
	fmt.Printf("Salario 01: %.2f\n", salario)
	salario, _ = calcSalario("B", 176)
	fmt.Printf("Salario 02: %.2f\n", salario)
	salario, _ = calcSalario("A", 172)
	fmt.Printf("Salario 03: %.2f\n", salario)
}
