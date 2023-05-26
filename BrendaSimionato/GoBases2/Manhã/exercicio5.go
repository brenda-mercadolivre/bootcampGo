package main

import (
	"errors"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogFunc(quantity int) int {
	return quantity * 10000
}

func catFunc(quantity int) int {
	return quantity * 5000
}

func hamsterFunc(quantity int) int {
	return quantity * 250
}

func tarantulaFunc(quantity int) int {
	return quantity * 150
}

func animal(t string) (func(quantity int) int, error) {
	if t == dog {
		return dogFunc, nil
	}
	if t == cat {
		return catFunc, nil
	}
	if t == hamster {
		return hamsterFunc, nil
	}
	if t == tarantula {
		return tarantulaFunc, nil
	}
	return nil, errors.New("Animal indefinido.")
}

func main() {
	// dfunc, _ := animal(dog)
	// fmt.Printf("quandidade de alimento do(s) cachorro(s) (em gramas):
	// %d gramas\n", dfunc(5))
	// cfunc, _ := animal(cat)
	// fmt.Printf("quandidade de alimento do(s) gato(s) (em gramas): %d
	// gramas\n", cfunc(8))
	// _, err := animal("invalid animal")
	// if err != nil {
	// fmt.Println(animal(dog, (5)), animal())
}

// }
