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

type erroCustomizadoSalario struct {
	horasTrabalhadas int
	mensagem         string
}

func (s *erroCustomizadoSalario) Error() string {
	return fmt.Sprintf("%d - %v", s.horasTrabalhadas, s.mensagem)
}

func calcSalario(horasTrabalhadas int, valorHora int) (int, error) {
	salario := (horasTrabalhadas * valorHora)
	return salario, &erroCustomizadoSalario{
		horasTrabalhadas: horasTrabalhadas,
		mensagem:         fmt.Sprintf("Salário totalizado no valor de: %d", salario),
	}
}

func validaImposto(salario float64) float64 {
	if salario >= 20.000 {
		desconto := (salario * 0.10)
		return salario - desconto
	} else {
		return salario
	}
}

func validaHoras(horasTrabalhadas int) (int, error) {
	if horasTrabalhadas <= 80 || horasTrabalhadas < 0 {
		return http.StatusBadRequest, &meuErroCustomizado{
			status:   http.StatusBadRequest,
			mensagem: "O trabalhador não pode ter trabalhado menos de 80 horas por mês",
		}
	} else {
		return http.StatusOK, &meuErroCustomizado{
			status:   http.StatusOK,
			mensagem: "Horas validadas com sucesso.",
		}
	}
}

func main() {
	salario, erro := calcSalario(100, 120)

	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	} else {
		validaHoras(salario)
		validaImposto(float64(salario))
		fmt.Printf("Salário no valor de: %d", salario)
	}
}
