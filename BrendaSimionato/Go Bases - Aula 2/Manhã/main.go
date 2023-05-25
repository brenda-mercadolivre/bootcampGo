package main

import "fmt"

func main() {
	//Exerc 1
	var name string = "Brenda"
	var age int = 25

	fmt.Println("Exercicio 1:", "Meu nome é", name, ", tenho", age, "anos.")

	//Exerc 2
	var temperatura int = 27
	var umidade int = 47
	var pressaoAtmosferica float32 = 1.1019

	fmt.Println("Exercicio 2:", "Temperatura atual:", temperatura, "%,umidade:", umidade, "%,pressão atmosférica:", pressaoAtmosferica)

	//Exerc 3
	var nome string = "Brenda"
	var sobrenome string = "Simionato"
	var idade int = 25
	var licencaParaDirigir = true
	var estaturaDaPessoa float32 = 1.58
	var quantidadeDeFilhos int = 0

	fmt.Println("Exercicio 3:", nome, sobrenome, idade, licencaParaDirigir, estaturaDaPessoa, quantidadeDeFilhos)

	//Exerc 4

	var sobrenome2 string = "Silva"
	var idade2 int = 32
	var filhos = false
	var salario float32 = 4585.90
	var nome2 string = "Fellipe"

	fmt.Println("Exercicio 4:", nome2, sobrenome2, idade2, filhos, salario)
}
