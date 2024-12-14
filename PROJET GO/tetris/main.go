package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	// on a besoin d'un texte d'entrée
	if len(args) != 1 {
		fmt.Println("on ne peut que lire dans un fichier")
		return
	}

	input, _ := os.ReadFile(args[0]) // input est de type byte
	if len(input) == 0 {
		fmt.Println("fichier vide")
		return
	}
	//fmt.Println(string(input))
	// Divise la chaîne en un tableau de tétrominos
	tetrominoString := strings.Split(string(input), "\n\n")

	// Affiche le nombre de tétrominos trouvés
	//fmt.Println("Nombre de tétrominos : ", len(tetrominoString))
	a := createTetrominos(tetrominoString)
	fmt.Println((a))

}

type Tetromino struct {
	shape [][]bool
}

// Crée un tableau de tétrominos à partir du tableau de chaînes
func createTetrominos(tetrominoStrings []string) []Tetromino {
	tetrominos := make([]Tetromino, len(tetrominoStrings))
	for i, tetrominoStr := range tetrominoStrings {
		shape := make([][]bool, 4)
		for j := range shape {
			shape[j] = make([]bool, 4)
		}
		for j, rowString := range strings.Split(tetrominoStr, "\n") { // on transforme tetrominoStr type rune en string
			for k, char := range rowString {
				if char == '#' {
					shape[j][k] = true
				}
			}
		}
		tetrominos[i] = Tetromino{shape}
	}
	return tetrominos
}

// Fonction qui détermine si un tétromino peut être placé dans une position donnée
func canPlace(tetromino [][]string, row int, col int, orientation int, square [][]string) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// Coordonnées du carré dans le tétromino
			tetrominoRow, tetrominoCol := i, j

			// Applique la rotation à la coordonnée du carré dans le tétromino
			for k := 0; k < orientation; k++ {
				tetrominoRow, tetrominoCol = 3-tetrominoCol, tetrominoRow
			}

			// Coordonnées du carré dans la grille
			squareRow, squareCol := row+i, col+j

			// Vérifie si la position est en dehors de la grille
			if squareRow < 0 || squareRow >= len(square) || squareCol < 0 || squareCol >= len(square[0]) {
				return false
			}

			// Vérifie si la position est déjà occupée par un autre carré
			if square[squareRow][squareCol] != "" && tetromino[tetrominoRow][tetrominoCol] != "" {
				return false
			}
		}
	}
	return true
}

// Fonction qui trouve la position et l'orientation qui minimise la taille de la grille

