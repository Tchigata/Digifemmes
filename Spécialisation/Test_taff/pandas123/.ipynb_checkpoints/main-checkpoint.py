import pandas as pd
import numpy as np

############################ Création d'un tableau NumPy #################################################
# Créer un tableau NumPy avec dtype=object
data_numpy = np.array([
    ['Bleu', [1, 2], 1.1],
    ['Rouge', [3, 4], 2.2],
    ['Rose', [5, 6], 3.3],
    ['Gris', [7, 8], 4.4],
    ['Noir', [9, 10], 5.5]
], dtype=object)

# Créer le DataFrame
df_numpy = pd.DataFrame(data_numpy, columns=['couleur', 'liste', 'nombre'])

# Afficher le DataFrame
print("DataFrame à partir d'un tableau NumPy :")
print(df_numpy)
# Afficher les types de chaque colonne pour le DataFrame NumPy
print("\nTypes des colonnes (NumPy) :")
print(df_numpy.dtypes)
print("\nTypes de la première valeur de chaque colonne (NumPy) :")
print(df_numpy.iloc[0].apply(type))


############################ Création d'une serie pandas #################################################
# Créer des séries pour chaque colonne
couleur = pd.Series(['Bleu', 'Rouge', 'Rose', 'Gris', 'Noir'])
liste = pd.Series([[1, 2], [3, 4], [5, 6], [7, 8], [9, 10]])
nombre = pd.Series([1.1, 2.2, 3.3, 4.4, 5.5])

# Créer le DataFrame
df_series = pd.DataFrame({'couleur': couleur, 'liste': liste, 'nombre': nombre})

# Afficher le DataFrame
print("\nDataFrame à partir de séries Pandas :")
print(df_series)

# Afficher les types de chaque colonne pour le DataFrame de séries
print("\nTypes des colonnes (Séries) :")
print(df_series.dtypes)
print("\nTypes de la première valeur de chaque colonne (Séries) :")
print(df_series.iloc[0].apply(type))