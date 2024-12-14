package gestionsPkg

import (
	"net/http"
	"os"
	"strings"
)

func AsciiArtPage(mots, police string) (string, int) {

	// Lecture du fichier de police
	mot, err := os.ReadFile("./ascii-art/banners/" + police + ".txt")
	if err != nil {
		return "Erreur lors de la lecture du fichier de police", http.StatusInternalServerError
	}

	contenu := strings.Split(string(mot), "\n")
	tab := strings.Split(strings.ReplaceAll(mots, "\r\n", "\\n"), "\\n")

	var resultat string
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
			resultat += resul + "\n"
		}
	}
	return resultat, http.StatusOK
}
