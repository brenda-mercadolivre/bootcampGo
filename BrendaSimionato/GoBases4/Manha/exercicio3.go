package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func salarioErroCustomizado(salario float64, status int) {
	if salario < 15.000 {
		fmt.Println(errors.New("O salário digitado não está dentro do valor mínimo."), salario)

		status := fmt.Errorf("\nmomento em que o erro aconteceu: %v", time.Now())
		fmt.Println("erro: ", status)
		return
	} else {
		fmt.Println(errors.New("Necessário pagamento de imposto."), salario)

		status := fmt.Errorf("\nmomento em que o erro aconteceu: %v", time.Now())
		fmt.Println("erro: ", status)
		return
	}
}

func main() {
	salarioErroCustomizado(2.000, http.StatusBadRequest)
}
