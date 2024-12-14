package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// creation de la variable qui recupera la valeure du flag
var (
	nameFile string
)

// recuperation de la flag
func RecupFlag() string {
	// recuperation de la flag avec la fonction StringVar()
	flag.StringVar(&nameFile, "color", "", "le nom de la couleur")

	// parser
	flag.Parse()
	return nameFile
}

func RecupArgs() int {
	result := len(flag.Args())

	return result
}

func main() {

	// la couleur
	var coloration = RecupFlag()

	// recuperation des arguments
	args := os.Args[1:]
	arg_donnee := RecupArgs()

	// verification du nombre argument donnees
	if (len(args) == 2 && arg_donnee != 1) || (len(args) == 3 && arg_donnee != 2) || (len(args) > 3 && arg_donnee >= 2) {
		erreur := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\""
		fmt.Println(erreur)
	} else {

		// variables
		var verif_chaine_color bool

		// declaration des couleurs dans la map coleur
		colorReset := "\033[0m"

		coleur := make(map[string]string)

		coleur["red"] = "\033[31m"
		coleur["green"] = "\033[32m"
		coleur["yellow"] = "\033[33m"
		coleur["purple"] = "\033[35m"
		coleur["blue"] = "\033[34m"
		coleur["cyan"] = "\033[36m"
		coleur["white"] = "\033[37m"
		coleur["black"] = "\033[30m"
		coleur["teal"] = "\033[36m"
		coleur["orange"] = "\033[38;5;208m"
		coleur["pink"] = "\033[38;5;200m"

		// verifie si il y a un element specifique a colorer
		if len(args) == 3 {
			verif_chaine_color = true

		}
		// ouverture du fichier banner
		file, err := os.Open("standard.txt")

		// verification d'erreur
		if err != nil {
			fmt.Println("erreur")
		}

		if verif_chaine_color {

			// recuperation de la chaine a colorer
			chaineColor := args[1]

			// recuperation de la chaine donnee
			chaine := args[2]
			//fmt.Println(chaine)

			// verifie si la chaine est un retour a la ligne
			if chaine == "\\n" {
				fmt.Print("\n")
			} else { // sinon

				// mes variables
				var tab_ascii []string

				// scanne du fichier banner
				scane := bufio.NewScanner(file)

				// scanne du fichier par ligne et ajout dans un tableau tab_ascii
				for scane.Scan() {
					line := scane.Text()
					tab_ascii = append(tab_ascii, string(line))
				}

				// recuperation de la chaine entree et conversion en tableau
				tab_chaine := strings.Split(chaine, "\\n")

				// Recuperation et Ecriture de l'ascii art

				// pour le B RGB()
				if len(chaineColor) == 1 { // permet de colorer un caractere specifique dans la chaine

					for _, chain := range tab_chaine {

						fmt.Print(string(colorReset))           // supprime la couleur en tampon
						rune_chaineColor := []rune(chaineColor) // tableau de rune de la chaine a colorer

						if strings.Contains(chain, chaineColor) { // verifie si la chaine a afficher contient la chaine a colorer

							for i := 0; i < 8; i++ {

								ligne := 1 + (int(rune_chaineColor[0]-32)*9 + i) // recuperation de la ligne

								// affiche les caracteres
								for _, valrun := range chain {
									if valrun == rune_chaineColor[0] { // verifie le caractere a colorer
										// color le caractere correspondant
										fmt.Print(string(coleur[coloration]), tab_ascii[ligne]) // affiche les premieres ligne de chaque caractere
										fmt.Print(string(colorReset))                           // supprime la couleur en tampon
									} else {
										ligne := 1 + (int(valrun-32)*9 + i) // recupere la ligne
										fmt.Print(tab_ascii[ligne])         // affiche les premieres ligne de chaque caractere
									}

								}

								// si la chain n'est pas vide affiche un retour a la ligne
								if chain != "" {
									fmt.Println()
								}
							}

							// si la chain est vide affiche un retour a la ligne
							if chain == "" {
								fmt.Println()
							}
						}
					}
					fmt.Print(string(colorReset)) // supprime la couleur en tampon

				} else { // permet de colorer un mot specifique dans la chaine donnee // Pour hey Guys

					for _, chain := range tab_chaine {

						fmt.Print(string(colorReset)) // supprime la couleur en tampon

						tab_chain_color := strings.Split(chaineColor, " ") // convertie la chaine a colorer en tableau
						var index int
						var last_index int

						if len(tab_chain_color) > 1 { // verifie si la chaine a colorer comprends plusieurs mots

							index = strings.Index(chain, tab_chain_color[0])
							// recupere l'index de debut de la chaine a colorer
							last_index = strings.Index(chain, tab_chain_color[len(tab_chain_color)-1]) + len(tab_chain_color) // recupere l'index de fin de la chaine a colorer
						} else {
							index = strings.Index(chain, tab_chain_color[0])                                    // recupere l'index de debut de la chaine a colorer
							last_index = strings.Index(chain, tab_chain_color[0]) + len(tab_chain_color[0]) - 1 // recupere l'index de fin de la chaine a colorer

						}
						fmt.Println(tab_chain_color[0])

						if strings.Contains(chain, chaineColor) { // verifie si la chaine a afficher contient la chaine a colorer

							for i := 0; i < 8; i++ {

								// affiche les caracteres
								for j, valrun := range chain {

									if j >= index && j <= last_index { // verifie si l'index du mot a colorer commence

										fmt.Print(string(colorReset))       // supprime la couleur en tampon
										ligne := 1 + (int(valrun-32)*9 + i) // recupere la ligne

										// color le caractere correspondant
										fmt.Print(string(coleur[coloration]), tab_ascii[ligne]) // affiche les premieres ligne de chaque caractere
										fmt.Print(string(colorReset))                           // supprime la couleur en tampon

									} else {
										fmt.Print(string(colorReset))       // supprime la couleur en tampon
										ligne := 1 + (int(valrun-32)*9 + i) // recupere la ligne
										fmt.Print(tab_ascii[ligne])         // affiche les premieres ligne de chaque caractere
									}

								}

								// si la chain n'est pas vide affiche un retour a la ligne
								if chain != "" {
									fmt.Println()
								}
							}

							// si la chain est vide affiche un retour a la ligne
							if chain == "" {
								fmt.Println()
							}
						}
					}

				}

			}

		} /* else { // permet de colorer toute la chaine quand aucun element specifique est donné

			// recuperation de la chaine donnee
			chaine := args[1]

			// verifie si la chaine est un retour a la ligne
			if chaine == "\\n" {
				fmt.Print("\n")
			} else {
				fmt.Print(string(colorReset)) // supprime la couleur en tampon

				// mes variables
				var tab_ascii []string

				// scanne du fichier banner
				scane := bufio.NewScanner(file)

				// scanne du fichier par ligne et ajout dans un tableau tab_ascii
				for scane.Scan() {
					line := scane.Text()
					tab_ascii = append(tab_ascii, string(line))
				}

				// recuperation de la chaine entree et conversion en tableau
				tab_chaine := strings.Split(chaine, "\\n")

				// Recuperation et Ecriture de l'ascii art
				for _, chain := range tab_chaine {
					for i := 0; i < 8; i++ {
						for _, valrun := range chain {

							ligne := 1 + (int(valrun-32)*9 + i)                     // recupere la ligne
							fmt.Print(string(coleur[coloration]), tab_ascii[ligne]) // affiche les premieres ligne de chaque caractere et applique la coloration

						}
						// si la chain n'est pas vide affiche un retour a la ligne
						if chain != "" {
							fmt.Println()
						}
					}

					// si la chain est vide affiche un retour a la ligne
					if chain == "" {
						fmt.Println()
					}

				}
				fmt.Print(string(colorReset)) // supprime la couleur en tampon
			}

		}*/

	}

}
