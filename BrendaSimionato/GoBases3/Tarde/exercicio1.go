package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

func (u usuario) detalhar() {
	fmt.Println("Nome:", u.nome)
	fmt.Println("Sobrenome:", u.sobrenome)
	fmt.Println("Idade:", u.idade)
	fmt.Println("Email:", u.email)
	fmt.Println("Senha:", u.senha)
}

func main() {
	u := usuario{
		nome:      "Brenda",
		sobrenome: "Simionato",
		idade:     25,
		email:     "brenda.simionato@mercadolivre.com",
		senha:     "ausente",
	}
	u.detalhar()

	var nome, sobrenome, email, senha string
	var idade int

	fmt.Println("Digite o seu nome: ")
	fmt.Scan(&nome)
	fmt.Println("Digite o seu sobrenome: ")
	fmt.Scan(&sobrenome)
	fmt.Println("Digite a sua idade: ")
	fmt.Scan(&idade)
	fmt.Println("Digite o seu email: ")
	fmt.Scan(&email)
	fmt.Println("Digite a sua senha: ")
	fmt.Scan(&senha)

}
