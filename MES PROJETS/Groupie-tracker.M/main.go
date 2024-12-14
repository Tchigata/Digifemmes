package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

// creer deux pages, une page pour avoir la liste des artistes et les autres liens pour afficher leurs details
// artist
// pour utiliser un fichier json il faut le composer en structure
// http.get pour lire un fichier en ligne
// ioutil.ReadAll pour lire le fichier
type artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	Creationdate int
	Firstalbum   string
	Locations    string
	Concertdates string
	Relations    string
}

// location
type location struct {
	Id        int
	Locations []string
	Dates     string
}

// dates
type dates struct {
	Id    int
	Dates []string
}

// relation
type relation struct {
	Id             int
	DatesLocations map[string][]string
}

type artistReturn struct {
	Artists []artist
}

type detailReturn struct {
	Artists   artist
	Locations location
	Date      dates
	Relations relation
	Title     string
}

func getArtist() ([]artist, error) {

	response, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	bytes, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var productsList []artist
	//Unmrsahll est une fonction qui permet de transformé les données js recupéré en d'autre type de fonction
	err := json.Unmarshal(bytes, &productsList)

	if err != nil {
		return nil, err
	}

	return productsList, nil
}

func Home(w http.ResponseWriter, r *http.Request) {

	artists, err := getArtist()

	if err != nil {
		panic(err)
	}

	view, _ := template.ParseFiles("myfiles/index.html")

	data := artistReturn{Artists: artists}

	view.Execute(w, data)

}
func Detail(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	// Extract specific parameter values
	param1Value := queryValues.Get("id")

	response, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + param1Value)
	bytes, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	var relationsList relation
	json.Unmarshal(bytes, &relationsList)

	response2, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + param1Value)
	bytes2, _ := ioutil.ReadAll(response2.Body)
	response2.Body.Close()
	var locationsList location
	json.Unmarshal(bytes2, &locationsList)

	response3, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + param1Value)

	bytes3, _ := ioutil.ReadAll(response3.Body)

	response3.Body.Close()
	var artistList artist

	json.Unmarshal(bytes3, &artistList)

	response4, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + param1Value)
	bytes4, _ := ioutil.ReadAll(response4.Body)
	response4.Body.Close()

	var datesList dates
	json.Unmarshal(bytes4, &datesList)

	view, _ := template.ParseFiles("myfiles/detail.html")

	data := detailReturn{Artists: artistList, Locations: locationsList, Date: datesList, Relations: relationsList, Title: param1Value}

	view.Execute(w, data)

}
func main() {

	styles := http.FileServer(http.Dir("./myfiles/static"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	http.HandleFunc("/", Home)
	http.HandleFunc("/artist", Home)
	http.HandleFunc("/detail", Detail)
	fmt.Println("http://localhost:8080/")
	OuvrirNavigateur("http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

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
