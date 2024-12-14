package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var grille [][]string

func main() {
	args := os.Args[1]
	// on a besoin d'un texte d'entrée
	file, err := os.Open(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	resultat := Input(file)
	Horizon(resultat)
	print()
}

func Input(file io.Reader) [4][4]string {
	//var tetrominoArray [][4][4]string
	var tetromino [4][4]string
	scanner := bufio.NewScanner(file)
	i, in, flag, alpha := 0, 0, true, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			if flag {
				flag = false
				continue
			} else {
				fmt.Println("ERROR")
				os.Exit(0)
			}
		}
		var arr [4]string // c'est pour gerer la ligne car le nombre d'element doit egale 4
		// arr répresente le tableau me permettant de parcourir ligne par ligne
		if len(str) != 4 {
			fmt.Println("ERROR")
			os.Exit(0)
		}
		for ind, _ := range arr { // c' est pour identifier la forme du tétrominos
			if rune(str[ind]) == '.' {
				arr[ind] = "."
			} else if rune(str[ind]) == '#' {
				arr[ind] = string(alpha[in])
			} else {
				fmt.Println("ERROR")
				os.Exit(0)
			}
		}
		tetromino[i] = arr // c' est pour remplacer la forme du tétrominos par les lettres de l'alphabet
		i++
		if i == 4 {
			flag = true
			i = 0
			in++
		}
	}
	return tetromino

}

// cette fonction verifie que le tétrominos est adjacent et qu'il contient 4 élément
func Forma(tetromino [4][4]string) bool {
	c, d := 0, 0

	for a, elem := range tetromino { // a lit les lignes et b les colonnes
		for b, elem2 := range elem {
			if elem2 != "." {
				d++
				if a+1 < 4 && tetromino[a+1][b] != "." { // je verifie l'adjacent en ligne c'est à dire vers le bas
					c++
				}
				if a-1 >= 0 && tetromino[a-1][b] != "." { //  je verifie l' adjacent en ligne c'est à dire vers le haut
					c++
				}
				if b+1 < 4 && tetromino[a][b+1] != "." { // verifie l' adjacent en colonne c'est à dire vers la droite
					c++
				}
				if b-1 >= 0 && tetromino[a][b-1] != "." { // verifie l'adjacent en colonne c'est à dire vers la droite
					c++
				}
			}
		}
	}
	if d == 4 {
		return true
	}
	if c == 6 || c == 8 { // le nombre de l'adjacent par defaut soit c= 6 ou c = 8
		return true
	}
	return false
}

/* fonction permettant de mettre le tétrominos à la premiere place pour completer le carré*/

// De manière vertical
func Horizon(tetromino [4][4]string) [4][4]string {
	//shifts tetromino row by 1
	temp := tetromino[0]
	tetromino[0] = tetromino[1]
	tetromino[1] = tetromino[2]
	tetromino[2] = tetromino[3]
	tetromino[3] = temp
	return tetromino
}

//Fonction pour afficher le resultat
func print() {
	for i := range grille {
		for j := range grille {
			fmt.Printf("%s", grille[i][j])
		}
		fmt.Printf("\n")
	}
}
