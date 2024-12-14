package gestionsPkg

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type Page struct {
	Interieur string
	Exterieur string
}

var sortie string
var statut int

// La  fonction nommée "RenduTemplate" prend en entrée un http.ResponseWriter et un http.Request.
// Elle est utilisée comme gestionnaire de requêtes HTTP pour un serveur web.
// Cette fonction s'occupe de la logique de traitement des requêtes GET et POST sur la racine du site ("/").
func Accueil(w http.ResponseWriter, dem *http.Request) {
	if dem.URL.Path == "/" { //Cette ligne vérifie si le chemin d'accès (Path) de la requête reçue est égal à "/". Si c'est le cas, cela signifie que la requête a été envoyée à la racine du site.
		switch dem.Method { //Cette ligne vérifie si la méthode de la requête est une méthode GET.

		case "POST":
			dem.ParseForm()
			if !CaractereValide(dem.Form.Get("input")) {
				MauvaiseRequete(w, dem)
			} else {
				sortie, statut = AsciiArtPage(dem.Form["input"][0], dem.Form["banniere"][0])
				if statut == 500 {
					StatusInternalServerError(w, dem)
				} else {
					exemple := Page{
						Interieur: dem.Form["input"][0],
						Exterieur: sortie,
					}
					templates, erreur := template.ParseFiles("templates/accueil.html")
					if erreur != nil { // Sinon, le modèle est exécuté avec les données nil en utilisant la fonction templates.Execute, et le résultat est renvoyé en réponse à la requête.
						http.Error(w, erreur.Error(), http.StatusInternalServerError)
						return
					}
					templates.Execute(w, exemple)

					/*
						Ces lignes appellent la fonction AsciiOutput pour générer une chaîne ASCII en utilisant les entrées de formulaire "input" et "bannier". Si la fonction SortieAscii renvoie une réponse HTTP 500, la fonction InternalServerError est appelée pour renvoyer une réponse HTTP 500. Sinon, la fonction crée une
					*/
				}
			}

		case "GET": //Si c'est le cas, cela signifie que la requête est une demande de récupération de la ressource à la racine du site.
			templates, err := template.ParseFiles("templates/accueil.html") //charge le fichier de modèle "templates/index.html" en utilisant la fonction template.ParseFiles.
			//Si une erreur se produit lors du chargement, la fonction StatusInternalServerError est appelée pour renvoyer une réponse HTTP 500.
			if err != nil { // Sinon, le modèle est exécuté avec les données nil en utilisant la fonction templates.Execute, et le résultat est renvoyé en réponse à la requête.
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			templates.Execute(w, nil)

			/*
			   Ces lignes analysent les données envoyées avec la requête POST en utilisant la fonction dem.ParseForm.
			   Ensuite, ils vérifient si le champ "input" du formulaire contient des caractères non-ASCII en utilisant
			   la fonction CaractereValide. Si c'est le cas, la fonction RequeteInvalide est appelée pour renvoyer une réponse HTTP 400 (mauvaise requête).
			   Sinon, la fonction continue avec la logique de traitement de la requête.
			*/

		}
	} else {
		StatutIntrouvable(w, dem)
	}
	erreur := os.WriteFile("fileascii.txt", []byte(sortie), 168)
	var err error
	if err != nil {
		panic(erreur)
	}
}

/*
Nous itérons sur les runes de la chaîne s. Nous faisons une comparaison sur la valeurc de chaque
rune avec 127 pour vérifier si elle dépasse la plage des caractères ASCII valide. si c'est le cas,
nous renvoyons 'false'. Sinon 'true' si tous les caractère de la chaîne sont des caractères ASCII valides.
*/
func CaractereValide(s string) bool {
	for _, chaine := range s {
		if chaine > 127 {
			return false
		}
	}
	return true
}

/*
Cette fonction est utilisée comme gestionnaire d'erreur pour renvoyer une réponse 400 Bad Request lorsqu'une requête contient des paramètres non valides.

La fonction commence par appeler la méthode WriteHeader de l'objet http.ResponseWriter pour définir le code de statut de la réponse à 400, ce qui indique que la requête est malformée ou ne peut pas être traitée en raison d'une erreur du client.

Ensuite, la fonction tente de charger le fichier HTML de la page d'erreur 400 à l'aide de la fonction template.ParseFiles. Si une erreur se produit lors du chargement du fichier, la fonction appelle la fonction InternalServerError pour renvoyer une réponse 500 Internal Server Error à la place.

Enfin, si le fichier est chargé avec succès, la fonction appelle la méthode Execute de l'objet de modèle pour générer la page HTML et la renvoyer en tant que réponse à la requête HTTP en cours.

En résumé, cette fonction est utilisée pour renvoyer une réponse 400 Bad Request avec une page HTML d'erreur personnalisée lorsque les paramètres de la requête sont invalides.

Rolande
func InternalServerError(w http.ResponseWriter, r *http.Request) {
*/
func MauvaiseRequete(w http.ResponseWriter, dem *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	templates, erreur := template.ParseFiles("templates/400.html")
	if erreur != nil { // Sinon, le modèle est exécuté avec les données nil en utilisant la fonction templates.Execute, et le résultat est renvoyé en réponse à la requête.
		http.Error(w, erreur.Error(), http.StatusInternalServerError)
		return
	}
	templates.Execute(w, nil)
}

/*
Cette fonction est appelée lorsqu'une erreur interne survient dans le serveur web. Elle prend en entrée un http.ResponseWriter et un http.Request et utilise le Writer pour envoyer une réponse au client avec un code d'état HTTP 500 (Internal Server Error).

Ensuite, elle charge le fichier de modèle "500.html" à partir du dossier "templates" et tente d'exécuter ce modèle. Si une erreur survient lors de l'exécution du modèle, elle est enregistrée à l'aide de la fonction log.Fatal(). Si tout se passe bien, le contenu généré par le modèle est renvoyé au client via le Writer.
*/
func StatusInternalServerError(w http.ResponseWriter, dem *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	t, _ := template.ParseFiles("templates/500.html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	Cette fonction suit le même principe que la fonction InternalServerError mais avec un code d'état HTTP 404 (Not Found)
	pour indiquer que la ressource demandée n'a pas été trouvée sur le serveur.

Elle prend également en entrée un http.ResponseWriter et un http.Request, utilise le Writer pour
envoyer une réponse au client avec un code d'état HTTP 404 (Not Found). Elle charge le fichier de modèle "404.html"
à partir du dossier "templates" et tente d'exécuter ce modèle. Si une erreur survient lors de l'exécution du modèle,
elle est enregistrée à l'aide de la fonction log.Fatal(). Si tout se passe bien, le contenu généré par le modèle est renvoyé au client via le Writer.
*/
func StatutIntrouvable(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		StatusInternalServerError(w, r)
		return
	}
	t.Execute(w, nil)
}

// La fonction appelée "OuvrirNavigateur" prend une URL (une chaîne de caractères) en argument et essaie
// d'ouvrir par défaut cette URL dans le navigateur en fonction du système d'exploitation en utilisant des commandes système.
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
