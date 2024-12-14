package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

const ort = ":1234"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A DIGIFEMME, WOMEN CAN DO MORE ")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contactez moi à :")
}

func main() {
	http.HandleFunc("/ascii", Home)
	http.HandleFunc("/contact", contact)
	OuvrirNavigateur("http://localhost:1234/ascii")
	http.ListenAndServe(ort, nil)
}

func OuvrirNavigateur(lienUrl string) {
	var erreur error //Utiliser pour stocker les erreurs qui peuvent survenir lors de l'exécution de commande système.
	switch runtime.GOOS {
	case "linux": //Ce block de code sera exécuté si le système d'exploitation est un Linux
		erreur = exec.Command("xdg-open", lienUrl).Start()
	case "windows": //Ce block de code sera exécuté si le système d'exploitation est un Windows
		erreur = exec.Command("rundll32", "url.dll,FileProtocolHandler", lienUrl).Start()
	case "darwin": //Ce block de code sera exécuté si le système d'exploitation est un macOS
		erreur = exec.Command("open", lienUrl).Start()
	}
	if erreur != nil {
		log.Fatal(erreur)
	}
}
