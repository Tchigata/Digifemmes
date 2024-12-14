# ascii-art-web

 # Présentation des projets :
 ascii-art est un programme qui prend une chaîne de caractères comme argument; il retourne une représentation graphique des caractères de la chaîne donnée. 
 Ascii-art web est la suite de ascii-art destiné à permettre l’utilisation en ligne du programme de base sur le web.
 
 # Comment l’utiliser :
 --Cloner le projet 
 Pour exécuter le code, vous devez demarrer le serveur en tapant dans le terminal "go run ."

 -- Cette commande ouvrira la page d' accueil dans votre navigateur sur l'adresse 
 "http://localhost:4000"

# Ascii-art-web :
--Vous pouvez utiliser trois polices différentes : "Standard", "Shadow" et "Thinkertoy".

--Gestion des erreurs :

-400 : Mauvaise demande
-404 : Page introuvable, fichier introuvable
-500 : Erreur du serveur interne

# Le projet a été réalisé par:

--Kouakou Akissi Rolande (akikouakou)
--Sarah Kouamé (sakouame)
--Coulibaly Tchewa Mariam (tcouliba)

# Détails d'implémentation : algorithme
Ce projet a été dévéloppé sous un algorithme de programmation modulaire. Vous trouverez à l'intérieur du dossier Ascii-art-web, trois dossiers, dont:
--ascii-art : qui est composé des sous-dossiers banners(contient les trois bannières ou polices des caractères ascii imprimable) et css(pour le style de nos fichiers html)
--gestionPkg : composé des fichiers ascii-art.go permettant d'imprimer en caractère ascii et gestion_func.go pour les fonctions du serveur implémenté.
--templates : qui contient la page principale accueil.html et les pages statut http server pour les érreurs coté serveur.