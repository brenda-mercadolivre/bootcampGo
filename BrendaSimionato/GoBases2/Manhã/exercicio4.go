package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMin(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Não é possível realizar o cálculo.")
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min, nil
}

func calcMax(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Não é possível realizar o cálculo.")
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, nil
}

func calcAvg(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Não é possível realizar o cálculo.")
	}
	avg := 0.0
	for _, v := range values {
		avg += v
	}
	return avg / float64(len(values)), nil
}

func getFunc(t string) (func(values ...float64) (float64, error),
	error) {
	if t == minimum {
		return calcMin, nil
	}
	if t == average {
		return calcAvg, nil
	}
	if t == maximum {
		return calcMax, nil
	}
	return nil, errors.New("Retorno de função inválido.")
}

func main() {
	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)
	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}
