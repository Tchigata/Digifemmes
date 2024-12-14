package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"jaune", "rouge"},
		{"vert", "mauve"},
		{"noir", "blanc"}, //la virgule est importante pour que le compilateur ne lance pas un message d'erreur
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "prada"
	b[0][1] = "lacoste"
	b[1][0] = "louboutin"
	b[1][1] = "gucci"
	b[2][0] = "dolce&gabanna"
	b[2][1] = "louis vuiton"
	fmt.Printf("\n")
	printarray(b)

	v := [3][2]string{}
	printarray(v)
}
