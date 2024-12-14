
let fruit = prompt("choisissez un fruit : ");

let userinput = fruit.toLowerCase();

switch(userinput) {
    case "banane" :
        alert("j'adore les bananes");
        break
    case "pomme" :
        alert("Vous avez choisi une pomme");
        break
    case "kiwi" :
        alert("les kiwi sont vertes")
        break
    default :
    alert("Je ne connais pas ce fruit")
}