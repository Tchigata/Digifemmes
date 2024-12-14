package main

import (
	"ascii-art-web/gestionsPkg"
	"fmt"
	"log"
	"net/http"
)

var port = ":2000"

func main() {
	http.HandleFunc("/", gestionsPkg.Accueil)
	//http.Handle("/ascii-art/", http.StripPrefix("/ascii-art/", http.FileServer(http.Dir("./ascii-art/"))))
	gestionsPkg.OuvrirNavigateur("http://localhost:2000")

	fmt.Println("http://localhost:2000")
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}

}
