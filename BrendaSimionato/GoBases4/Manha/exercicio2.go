package main

import (
	"errors"
	"fmt"
)

func salarioErroCustomizado(salario float64) {
	if salario < 15.000 {
		fmt.Println(errors.New("O salário digitado não está dentro do valor mínimo."), salario)
		return
	} else {
		fmt.Println(errors.New("Necessário pagamento de imposto."), salario)
		return
	}
}

func main() {
	salarioErroCustomizado(1.700)
}
