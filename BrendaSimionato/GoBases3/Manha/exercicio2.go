package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func exerc2() {
	arquivo, err := os.Open("produtosGerados.csv")
	if err != nil {
		log.Println(err)
	}
	defer arquivo.Close()
	w := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t',
		tabwriter.AlignRight)
	scanner := bufio.NewScanner(arquivo)
	scanner.Scan()
	cabecalho := strings.Split(scanner.Text(), ",")
	for _, c := range cabecalho {
		fmt.Fprintf(w, "%s\t", c)
	}
	fmt.Fprintln(w)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}
