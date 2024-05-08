# devinette-algo-genetic

Petit test sur l'algo genetic 
Vidéo insipîration : 

(amy plant - algo genetic sur snake)
https://www.youtube.com/watch?v=YRjq02WT_oI&t=1s


La population commence avec une suite de numéro 
[1, 9, 8, 7] // a definir la longueur exemple 4

La population génére des numéros
On a une méthode de vérification qui consiste à voir si le premier est supérieur au deuxieme 

Humain 1
[1, 9, 8, 7]

1 > 9 ? Non
9 > 8 ? Oui
8 > 7 ? Oui

Score = 2

Humain 2 [5, 3, 7, 4]

5 > 3 ? Oui
3 > 7 ? Non
7 > 4 ? Oui

Score = 2

Humain 3 [2, 6, 4, 8]

2 > 6 ? Non
6 > 4 ? Oui
4 > 8 ? Non

Score = 1



etc....

On garde les deux meilleurs parents de la populations (humain 1 et 2)

Ceux qui ont des plus gros scores on les fait se "reproduire"

Et ainsi on recrée une population de x en se basant sur les parents

Les enfants obtienne un score... Se reproduisent (mais grandissent avant hein)... Pour créer une nouvelle génération... X fois
