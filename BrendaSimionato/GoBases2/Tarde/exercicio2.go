package main

import "fmt"

type Product struct {
	TipoProduto string
	Nome        string
	Preco       float64
}

func (p *Product) CalcularCusto() float64 {
	switch p.TipoProduto {
	case "médio":
		return p.Preco * 1.03
	case "grande":
		return p.Preco*1.06 + 2500
	default:
		return p.Preco
	}
}

type Ecommerce interface {
	Total() float64
	Adicionar(novoProduto IProduto)
}

type IProduto interface {
	CalcularCusto() float64
}

type Loja struct {
	Produtos []IProduto
}

func (loja *Loja) Adicionar(novoProduto IProduto) {
	loja.Produtos = append(loja.Produtos, novoProduto)
}

func (loja *Loja) Total() float64 {
	total := 0.0

	for _, produto := range loja.Produtos {
		total += produto.CalcularCusto()
	}

	return total
}

func novoProduto(tipoProduto, nome string, preco float64) IProduto {
	return &Product{
		TipoProduto: tipoProduto,
		Nome:        nome,
		Preco:       preco,
	}
}

func novoLoja() Ecommerce {
	return &Loja{}
}

func main() {
	loja := novoLoja()

	produto1 := novoProduto("pequeno", "pão de queijo", 4.00)
	produto2 := novoProduto("médio", "bolo de cenoura", 25.00)
	produto3 := novoProduto("grande", "bolo de casamento", 1000.00)

	loja.Adicionar(produto1)
	loja.Adicionar(produto2)
	loja.Adicionar(produto3)

	total := loja.Total()

	fmt.Println("Preço total: ", total)
}
