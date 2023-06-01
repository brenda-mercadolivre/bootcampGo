package main

import (
	"fmt"
	"net/http"
	"os"
)

type MeuErroCustomizado struct {
	status   int
	mensagem string
}

func (m *MeuErroCustomizado) Error() string {
	return fmt.Sprintf("%d - %v", m.status, m.mensagem)
}

func SalarioErroCustomizado(salario float64) (int, error) {
	if salario < 15.000 {
		return http.StatusBadRequest, &MeuErroCustomizado{
			status:   http.StatusBadRequest,
			mensagem: "O salário digitado não está dentro do valor mínimo",
		}
	} else {
		return http.StatusOK, &MeuErroCustomizado{
			status:   http.StatusOK,
			mensagem: "Necessário pagamento de imposto",
		}
	}
}

func exerc1() {
	salario, erro := SalarioErroCustomizado(1.200)

	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}
	fmt.Printf("Salário no valor de: %d", salario)

}
