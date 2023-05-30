package main

import (
	"fmt"
	"net/http"
	"os"
)

type meuErroCustomizado struct {
	status   int
	mensagem string
}

func (m *meuErroCustomizado) Error() string {
	return fmt.Sprintf("%d - %v", m.status, m.mensagem)
}

func salarioErroCustomizado(salario float64) (int, error) {
	if salario < 15.000 {
		return http.StatusBadRequest, &meuErroCustomizado{
			status:   http.StatusBadRequest,
			mensagem: "O salário digitado não está dentro do valor mínimo",
		}
	} else {
		return http.StatusOK, &meuErroCustomizado{
			status:   http.StatusOK,
			mensagem: "Necessário pagamento de imposto",
		}
	}
}

func main() {
	salario, erro := salarioErroCustomizado(1.200)

	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}
	fmt.Printf("Salário no valor de: %d", salario)

}
