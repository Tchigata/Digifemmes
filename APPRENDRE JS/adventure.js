//Générer un nombre aléatoire

function guessName()  {
    let randomNr =Math.floor( Math.random() *11);
   let guess ;

   do {
    guess = prompt("Devine un nombre compris entre 1 et 10")
    console.log(guess, randomNr)

    if (randomNr > guess) {
        console.log("Trop petit");
    } else if (randomNr < guess) {
        console.log("Trop grand")
    }
   }
   while(guess =! randomNr);
}

guessName();