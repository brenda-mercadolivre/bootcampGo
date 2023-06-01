package main

import (
	"errors"
	"fmt"
)

func SalarioErrCustomizado(salario float64) {
	if salario < 15.000 {
		fmt.Println(errors.New("O salário digitado não está dentro do valor mínimo."), salario)
		return
	} else {
		fmt.Println(errors.New("Necessário pagamento de imposto."), salario)
		return
	}
}

func exerc2() {
	SalarioErrCustomizado(1.700)
}
