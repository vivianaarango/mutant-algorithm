package main

import (
	"fmt"
	"strings"
)

func isMutant(adn []string) bool {
	mutantADN := make([][]string, len(adn))

	for i := 0; i < len(adn); i++ {
		// looping through the slice to declare
		// slice of slices of length 3
		separatedADN := strings.Split(adn[i], "")
		mutantADN[i] = make([]string, len(separatedADN))

		// assigning values to each
		// slice of a slice
		for j := 0; j < len(separatedADN); j++ {
			mutantADN[i][j] = separatedADN[j]
		}

		//fmt.Println(mutantADN[i])
	}

	fmt.Println("aqui", validateHorizontal(mutantADN))

	return true
}

func validateHorizontal(mutantADN [][]string) bool {
	countValid := 0
	for i := 0; i < len(mutantADN); i++ {
		coincidences := 0
		currentValue := ""
		for j := 0; j < len(mutantADN[i])-1; j++ {
			if currentValue == "" {
				currentValue = mutantADN[i][j]
			}

			if currentValue == mutantADN[i][j] {
				coincidences++
			}

			if coincidences == 4 {
				countValid++
				coincidences = 0
			}

			if countValid > 1 {
				return false
			}
		}
	}

	return true
}

func isValidADN(rowADN string) bool {
	separatedADN := strings.Split(rowADN, "")
	for _, adn := range separatedADN {
		if adn == "A" || adn == "T" || adn == "C" || adn == "G" {
			continue
		}
		return false
	}

	return true
}

func main() {
	data := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	//data := []string{"AAAAAE"}
	result := isMutant(data)

	if result {
		fmt.Println("Result:", "Es mutante")
	} else {
		fmt.Println("Result:", "No es mutante")
	}
}
