package main

import (
	"fmt"
	"os"
)

type GeraId struct {
	id int
}

func (geraId *GeraId) geradorId() int {
	if geraId.id < 0 || geraId == nil {
		panic("Inclua apenas valores positivos.")
	}

	geraId.id++
	return geraId.id
}

func geraErro() {
	if err := recover(); err != nil {
		fmt.Println("Erro:", err)
	}
}

func main() {

	defer geraErro()

	idGerado := GeraId{}
	idGerado.id = -7
	fmt.Println(idGerado.geradorId())

	idGerado.id = 10
	fmt.Println(idGerado.geradorId())

	_, err := os.Open("customers.txt")
	if os.IsNotExist(err) {
		panic("Erro: O arquivo indicado não foi encontrado ou está danificado.")
	}
}
