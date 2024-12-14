package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func average(str []float64) float64 {
	somme := 0.0
	for _, v := range str {
		somme += v
	}
	return somme / float64(len(str))
}

func variance(str []float64) float64 {
	a := average(str)
	var b float64
	for i := range str {
		b += math.Pow((str[i] - a), 2)
	}
	return b / float64(len(str)-1)
}

func ecartType(str []float64) float64 {
	a := variance(str)
	c := math.Sqrt(a)
	return c
}

func Regression(x, y []float64) (float64, float64) {

	n1 := len(x)

	var somme_xi_yi, somme_xi, somme_yi, somme_xi2, x_bar, y_bar float64

	for i := 0; i < n1; i++ {

		somme_xi_yi += x[i] * y[i]
		somme_xi += x[i]
		somme_yi += y[i]
		somme_xi2 += x[i] * x[i]

	}

	x_bar = average(x)
	y_bar = average(y)

	n := float64(n1)

	a := (somme_xi_yi - n*x_bar*y_bar) / (somme_xi2 - n*x_bar*x_bar)

	b := y_bar - a*x_bar

	return a, b

}

func Range(numbers []float64) (int, int) {
	x := make([]float64, len(numbers))
	y := numbers

	for i := range x {
		x[i] = float64(i)
	}

	a, b := Regression(x, y)
	next := float64(len(numbers))
	pred := a + b*next

	res := make([]float64, len(numbers))
	for i := range res {
		res[i] = y[i] - (a + b*x[i])
	}

	sd := ecartType(res)

	lower := int(pred - sd)
	if lower < 1 {
		lower = 1
	}
	upper := int(pred + sd)

	return lower, upper
}

func main() {

	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			break
		}

		numbers = append(numbers, x)

		if len(numbers) > 1 {
			a, b := Range(numbers[:len(numbers)-1])
			fmt.Printf("%d %d\n", a, b)
		}
	}

}
