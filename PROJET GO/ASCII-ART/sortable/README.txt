#explication#

Ce code est une application JavaScript qui récupère des données sur les super-héros à partir d'une API et les affiche 
dans un tableau avec des fonctionnalités de recherche, de tri et de pagination. Voici une explication étape par étape du code :

La variable data est un tableau vide qui va contenir les données des super-héros.

La fonction init() est une fonction asynchrone qui effectue les étapes suivantes :

Elle utilise la fonction fetch() pour récupérer les données des super-héros à partir de l'URL fournie.
Elle utilise la méthode json() pour extraire les données JSON de la réponse.
Ensuite, elle itère sur chaque élément des données récupérées et crée un tableau hero contenant les informations pertinentes du super-héros. 
Ce tableau hero est ensuite ajouté au tableau data.
Enfin, elle appelle la fonction render(data) pour afficher les éléments sur la page.
La fonction render(heroes, page) est responsable de l'affichage des super-héros sur la page en fonction des paramètres heroes et page. 
Elle effectue les étapes suivantes :

Récupère les valeurs des champs de recherche et de tri depuis les éléments HTML correspondants.
Efface les lignes précédemment ajoutées dans le tableau.
Effectue une recherche sur les super-héros en fonction du texte saisi et du champ de tri sélectionné.
Gère la pagination en déterminant les indices de début et de fin pour afficher les super-héros correspondants à la page actuelle.
Met à jour les informations affichées sur le nombre de résultats trouvés.
Génère les boutons de pagination en fonction du nombre total de pages.

Pour chaque super-héros à afficher, elle appelle la fonction createRow(k, i) pour créer une nouvelle ligne dans le tableau HTML.
La fonction createRow(arr, index) crée une nouvelle ligne dans le tableau HTML en utilisant les données du super-héros passées en paramètre.
Elle itère sur chaque élément du tableau arr et crée une cellule <td> pour chaque élément. 
Si c'est la première cellule, elle crée une balise <img> avec l'URL de l'image du super-héros, sinon elle crée un nœud de texte avec la valeur de l'élément.
Enfin, elle ajoute un gestionnaire d'événements de clic à la ligne pour afficher les détails du super-héros correspondant.

Le code suivant l'ajout des lignes de tableau définit des gestionnaires d'événements pour chaque en-tête de colonne du tableau. 
Lorsque l'utilisateur clique sur un en-tête de colonne, il déclenche une fonction de tri correspondante (sortStrings, sortNumStr, sortNum) 
pour trier les données en fonction de la colonne cliquée et de l'ordre de tri. Ensuite, la fonction render(table, 1) 
est appelée pour afficher les données triées sur la première page.

Les dernières parties du code définissent des gestionnaires d'événements pour la recherche, 
les sélections de colonnes de tri et l'affichage des détails d'un super-héros lors d'un clic sur une ligne du tableau. La fonction `show
