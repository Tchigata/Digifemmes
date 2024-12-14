package main

import "fmt"

func main() {

	var mois = [6]string{"janvier", "fevrier", "mars", "mai", "juin", "Juillet"}
	for i := 0; i < len(mois); i++ {
		fmt.Println(mois[i], "est le", i+1, "mois de l'année")
	}
}
