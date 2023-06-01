package main

import (
	"fmt"
	"os"
)

func exerc1() {

	fmt.Println("Iniciando...")

	_, erro := os.Open("customers.txt")

	if erro != nil {
		panic(erro)
	}

	fmt.Println("O arquivo indicado não foi encontrado ou está danificado.")

}
