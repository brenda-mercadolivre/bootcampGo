package main

import (
	"fmt"
)

type Usuario struct {
	id        int
	nome      string
	sobrenome string
	email     string
	produto   []Produto
}

type Produto struct {
	nome       string
	preco      float64
	quantidade int
}

func novoProduto(nome string, preco float64) Produto {
	fmt.Println("Nome:", nome)
	fmt.Println("Preço:", preco)

	return Produto{
		nome:  nome,
		preco: preco,
	}
}

func addProduto(usuario Usuario, produto Produto, quantidade int) {
	produto.quantidade = quantidade
	usuario.produto = append(usuario.produto, produto)
}

// func main() {

// }
