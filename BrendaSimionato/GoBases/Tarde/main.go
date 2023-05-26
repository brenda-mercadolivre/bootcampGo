package main

import "fmt"

func main() {
	//Exerc 1

	name := "Brenda"
	fmt.Println(len(name))

	//Exerc 2
	idade := 29
	empregado := true
	tempoTrabalho := 2
	salario := 100.000

	if empregado && idade > 22 && tempoTrabalho > 1 {
		if salario >= 100.000 {
			fmt.Println("Possível a realização do empréstimo sem juros.")
		} else {
			fmt.Println("Possível a realização do empréstimo.")
		}
	} else {
		fmt.Println("Não é possível a realização do empréstimo.")
	}

	//Exerc 3

	dias := 2
	switch dias {
	case 1:
		fmt.Println("Janeiro")

	case 2:
		fmt.Println("Fevereiro")

	case 3:
		fmt.Println("Marco")

	case 4:
		fmt.Println("Abril")

	case 5:
		fmt.Println(" Maio")

	case 6:
		fmt.Println(" Junho")

	case 7:
		fmt.Println("Julho")

	case 8:
		fmt.Println("Agosto")

	case 9:
		fmt.Println("Setembro")

	case 10:
		fmt.Println("Outubro")

	case 11:
		fmt.Println("Novembro")

	case 12:
		fmt.Println("Dezembro")
	}

	//Exerc 4

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
