package main

import "fmt"

func calcMedia(notas ...float64) float64 {
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	return soma / float64(len(notas))
}
func main() {
	fmt.Printf("Calcula Média: %.2f\n", calcMedia(10, 8, 9))
}
