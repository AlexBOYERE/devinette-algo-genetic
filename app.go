package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Taille de la population pour l'AG
const populationSize = 10

// Taille de la liste à trier
// Jouer avec pour voir l'évolution + grand + de suite à trouver
const listLength = 10

// Taux de mutation (10%)
const mutationRate = 0.1

// Generation des enfants devenant parents
// Jouer avec pour voir l'évolution + grand + de chance de trouver une bonne suite
const generationSize = 50

// Compte les points, si la permutation i est supérieur à i+1 alors le parent à bien mis dans l'ordre
// Simple, non ?
// A la fin il retourne le score, + le score est haut mieux c'est
func fitness(permutation []int) int {
	score := 0
	for i := 0; i < len(permutation)-1; i++ {
		if permutation[i] > permutation[i+1] {
			score++
		}
	}

	return score
}

// Selectionne deux parents aléatoirement dans la populations
// Recommence si le parent est reproduit deux fois (attention l'auto consanguinité)
func selectParents(population [][]int) ([]int, []int) {
	// Selectionne deux parents aléatoires
	parent1Index := rand.Intn(populationSize)
	parent2Index := rand.Intn(populationSize)
	for parent2Index == parent1Index {
		parent2Index = rand.Intn(populationSize)
	}

	// Sélectionner le parent le plus compétent (celui qui a réussi le plus de suite)
	parent1 := population[parent1Index]
	parent2 := population[parent2Index]
	if fitness(parent1) < fitness(parent2) {
		parent1, parent2 = parent2, parent1
	}

	//fmt.Println("Les parents :", parent1, parent2)

	return parent1, parent2
}

// Pour but de faire reproduire deux parents pour créer un enfant
// Récupère un point de croisement aléatoire dans la permutation (len(parent1))
// L'enfant est en deux, la première partie est du parent dominant (le plus gros score) ; la deuxieme pour l'autre
// retourne l'enfant créé par le croisement
func crossover(parent1 []int, parent2 []int) []int {
	// Choisir un point de croisement aléatoire
	crossoverPoint := rand.Intn(len(parent1)-1) + 1

	// Combine les deux parents pour créer l'enfant
	child := make([]int, len(parent1))
	copy(child, parent1[:crossoverPoint])
	copy(child[crossoverPoint:], parent2[crossoverPoint:])

	//fmt.Println("L'enfant", child)

	return child
}

// Génère le croisement entre les deux parents
func mutate(permutation []int) []int {
	// Parcourir la permutation et muter aléatoirement les éléments
	for i := range permutation {
		if rand.Float64() < mutationRate {
			// Inverser l'ordre
			j := rand.Intn(len(permutation))
			for j == i {
				j = rand.Intn(len(permutation))
			}
			permutation[i], permutation[j] = permutation[j], permutation[i]
		}
	}

	fmt.Println("La permutation (mutation) de l'enfant", permutation)

	return permutation
}

func main() {
	// Initialiser le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Générer la population initiale
	population := make([][]int, populationSize)
	for i := range population {
		permutation := make([]int, listLength)
		for j := range permutation {
			permutation[j] = j
		}
		rand.Shuffle(len(permutation), func(i, j int) {
			permutation[i], permutation[j] = permutation[j], permutation[i]
		})
		population[i] = permutation
	}

	// Afficher la population initiale
	fmt.Println("Population initiale:")
	for _, permutation := range population {
		fmt.Println(permutation)
	}
	fmt.Println("Fin de la population initial")

	// Evlotion de la populasse
	for generation := 0; generation < generationSize; generation++ {
		// Sélectionner les parents
		parent1, parent2 := selectParents(population)

		// Effectuer le croisement et la mutation pour créer de nouveaux enfants
		newChildren := make([][]int, 0)
		for i := 0; i < populationSize; i++ {
			child := crossover(parent1, parent2)
			child = mutate(child)
			newChildren = append(newChildren, child)
		}

		// Le grand remplacement (ancienne generation par la nouvelle)
		population = newChildren

		// récupérer le meilleur civil de l'ancienne génération
		bestPermutation := population[0]
		for i := 1; i < populationSize; i++ {
			if fitness(population[i]) < fitness(bestPermutation) {
				bestPermutation = population[i]
			}
		}
		fmt.Println("\n\n", "\nGénération", generation+1, ":", bestPermutation, " (fitness:", fitness(bestPermutation), ")", "\n\n")
	}

	// Afficher la meilleure solution finale
	fmt.Println("\nSolution finale:", population[0], "(fitness:", fitness(population[0]), ")")
}
