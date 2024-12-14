package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	strStandard()
}

func strStandard() {
	mot, _ := os.ReadFile(os.Args[2] + ".txt")
	arg := os.Args[1]

	// cette partie permet de verifier le format de l'argument passée
	if len(os.Args) != 3 || os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	contenu := strings.Split(string(mot), "\n")
	tab := strings.Split(arg, "\\n")

	for _, ligne := range tab { //Cette ligne parcours le tableau de string output pour le ranger dans un string
		for i := 0; i < 8; i++ {
			resul := ""
			for _, caractAscii := range ligne {
				index := (caractAscii-32)*9 + 1
				resul += contenu[int(index)+i]
			}
			if ligne == "" {
				break
			}
			fmt.Println(resul)
		}
	}
}
