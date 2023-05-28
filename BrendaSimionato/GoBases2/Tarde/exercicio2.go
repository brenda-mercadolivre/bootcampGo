package main

type produto struct {
	tipo  string
	nome  string
	preço float64
}

func (p produto) produto() float64 {
	return p.preço
}

func (p produto) medio() float64 {
	return p.preço * 0.3
}

func (p produto) grande() float64 {
	return (p.preço * 0.6) + 2.500
}

type Produto interface {
	calcCusto() float64
}

func novoProduto(tipo string, nome string, preço float64) produto {
	return produto{preço: preço, tipo: tipo, nome: nome}
}

const (
	tipoPequeno = "pequeno"
	tipoMedio   = "medio"
	tipoGrande  = "grande"
)

// func novosProdutos(tipoProduto string, valor ...float64) Produto {
// 	switch tipoProduto {
// 	case tipoPequeno:
// 		return pequeno{preço: valor[0]}

// 	case tipoMedio:
// 		return medio{preço: valor[0]}

// 	case tipoGrande:
// 		return grande{preço: valor[0]}
// 	}
// 	return nil
// }

// func main() {

// 	p := novoProduto(tipoPequeno, "pequeno", 10)
// 	fmt.Println("Tipo Pequeno")
// 	fmt.Printf("Nome: %.2f\n", p.nome)
// 	fmt.Printf("\tPreço: %.2f\n\t", p.calcCusto())

// 	fmt.Println()

// 	m := novoProduto(tipoMedio, "medio", 20)
// 	fmt.Println("Tipo Médio")
// 	fmt.Printf("Nome: %.2f\n", m.nome)
// 	fmt.Printf("\tPreço: %.2f\n\t", m.calcCusto())

// 	fmt.Println()

// 	g := novoProduto(tipoGrande, "grande", 30)
// 	fmt.Println("Tipo Grande")
// 	fmt.Printf("Nome: %.2f\n", g.nome)
// 	fmt.Printf("\tPreço: %.2f\n\t", g.calcCusto())
// }
