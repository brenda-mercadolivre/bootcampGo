package main

import (
	"errors"
	"fmt"
	"os"
)

type produto struct {
	id         int
	quantidade int
	preco      float64
}

func (p produto) linhaCSV() string {
	return fmt.Sprintf("%d,%d,%.2f\n", p.id, p.quantidade, p.preco)
}
func (p produto) cabecalhoCSV() string {
	return "id,quantidade,preco\n"
}

func geradorCsv(caminho string, produtos []produto) error {
	if len(produtos) == 0 {
		return errors.New("Não existem produtos para serem guardados.")
	}
	arquivo, err := os.OpenFile(caminho, os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	p := produtos[0]
	if _, err = arquivo.WriteString(p.linhaCSV()); err != nil {
		return fmt.Errorf("Ocorreu um erro: %w", err)
	}

	return nil
}
func main() {
	produtos := []produto{
		{
			id:         12,
			quantidade: 11,
			preco:      49.09,
		},
		{
			id:         6,
			quantidade: 7,
			preco:      2.82,
		},
		{
			id:         97,
			quantidade: 6,
			preco:      15.00,
		},
	}
	geradorCsv("produtosGerados.csv", produtos)
}
