package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Format()
	AsciartColor()

}

var color string
var long_arg bool
var Colors = map[string]string{
	"Init":   "\033[0m",
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"Gray":   "\033[37m",
	"Orange": "\033[38;2;255;140;0m",
	"Pink":   "\033[38;2;255;105;108m",
	"Indigo": "\033[38;2;138;43;226m",
	"Brown":  "\033[38;2;165;42;42m",
}

func AsciartColor() {
	color = getColorCode(strings.ToLower(os.Args[1][8:]))
	text := os.Args[2]
	letterToColor := text
	if len(os.Args)-1 > 2 {
		text = os.Args[3]
		letterToColor = os.Args[2]
	}
	if len(os.Args) == 3 {
		long_arg = true
	}

	strStandard(text, letterToColor)

}

func strStandard(input, letterToColor string) {
	var resultat string
	mot, _ := os.ReadFile("standard.txt")
	var str []string
	contenu := strings.Split(string(mot), "\n") // le banner
	tab := strings.Split(input, "\\n")
	for _, data := range contenu {
		str = append(str, data)
	}
	// mettre en couleur une lettre bien spécifique
	if long_arg {
		for _, ligne := range tab { //Cette ligne parcours le tableau de string output pour le ranger dans un string
			for i := 0; i < 8; i++ {
				for _, caractAscii := range ligne {
					if int(caractAscii) >= 32 && int(caractAscii) <= 126 {
						index := (caractAscii-32)*9 + 1
						resul := ""

						if strings.ContainsRune(letterToColor, caractAscii) {
							resul = (color + str[int(index)+i] + Colors["Init"])

						} else {
							resul = str[int(index)+i]
						}

						if resul != "" {
							resultat += resul
						}

					} else {
						fmt.Println("Invalid Input")
						return
					}

				}
				resultat += "\n"
			}

			if ligne == "" {
				fmt.Println()
			} else {
				fmt.Print(resultat)
			}
		}
	} else {
		if len(letterToColor) == 1 {
			for _, ligne := range tab { //Cette ligne parcours le tableau de string output pour le ranger dans un string
				for i := 0; i < 8; i++ {
					for _, caractAscii := range ligne {
						if int(caractAscii) >= 32 && int(caractAscii) <= 126 {
							index := (caractAscii-32)*9 + 1
							resul := ""

							if strings.ContainsRune(letterToColor, caractAscii) {
								resul = (color + str[int(index)+i] + Colors["Init"])

							} else {
								resul = str[int(index)+i]
							}

							if resul != "" {
								resultat += resul
							}

						} else {
							fmt.Println("Invalid Input")
							return
						}

					}
					resultat += "\n"
				}

				if ligne == "" {
					fmt.Println()
				} else {
					fmt.Print(resultat)
				}
			}
		} else {

			for _, ligne := range tab { //Cette ligne parcours le tableau de string output pour le ranger dans un string

				tab_lettercolor := strings.Split(letterToColor, " ") // convertie la chaine a colorer en tableau
				var indexe int
				var last_index int

				if len(tab_lettercolor) > 1 { // verifie si la chaine a colorer comprends plusieurs mots

					indexe = strings.Index(ligne, tab_lettercolor[0])                                                 // recupere l'index de debut de la chaine a colorer
					last_index = strings.Index(ligne, tab_lettercolor[len(tab_lettercolor)-1]) + len(tab_lettercolor) // recupere l'index de fin de la chaine a colorer

				} else {
					indexe = strings.Index(ligne, tab_lettercolor[0])                                   // recupere l'index de debut de la chaine a colorer
					last_index = strings.Index(ligne, tab_lettercolor[0]) + len(tab_lettercolor[0]) - 1 // recupere l'index de fin de la chaine a colorer

				}

				if strings.Contains(ligne, letterToColor) {

					for i := 0; i < 8; i++ {
						for j, caractAscii := range ligne {

							if j >= indexe && j <= last_index {

								fmt.Print(Colors["Init"])
								index := 1 + (int(caractAscii-32)*9 + i)

								fmt.Print(color, str[index])
								fmt.Print(Colors["Init"])

							} else {
								fmt.Print(Colors["Init"])
								index := 1 + (int(caractAscii-32)*9 + i)
								fmt.Print(str[index])
							}

						}

						if ligne != "" {
							fmt.Println()
						}
					}

					if ligne != "" {
						fmt.Println()
					}

				}

			}

		}

	}

}

// gestion du format
func Format() {
	switch {
	case len(os.Args) > 4, strings.HasPrefix(os.Args[1], "--color=") == false:
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEx: go run . --color=<color> <letters to be colored> \"something\" ")
		os.Exit(0)
	}

}

// Fonction qui gere les couleurs
func getColorCode(color string) string {
	switch color {
	case "red":
		return Colors["Red"]
	case "green":
		return Colors["Green"]
	case "yellow":
		return Colors["Yellow"]
	case "blue":
		return Colors["Blue"]
	case "purple":
		return Colors["Purple"]
	case "cyan":
		return Colors["Cyan"]
	case "gray":
		return Colors["Gray"]
	case "orange":
		return Colors["Orange"]
	case "pink":
		return Colors["Pink"]
	case "indigo":
		return Colors["Indigo"]
	case "brown":
		return Colors["Brown"]
	default:
		log.Fatalf("Invalid Color !!")
		return Colors["Init"]
	}
}
