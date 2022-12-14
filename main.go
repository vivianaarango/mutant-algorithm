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

		fmt.Println(mutantADN[i])
	}

	fmt.Println("Vertical:", vertical(mutantADN))
	fmt.Println("Horizontal:", horizontal(mutantADN))

	if horizontal(mutantADN) > 1 || vertical(mutantADN) > 1 {
		return false
	}

	if horizontal(mutantADN) == 1 || vertical(mutantADN) == 1 {
		return true
	}

	return false
}

func horizontal(mutantADN [][]string) int {
	countValid := 0
	for r := 0; r < len(mutantADN); r++ {
		coincidences := 1
		currentValue := ""

		for c := 0; c < len(mutantADN[r])-1; c++ {
			if currentValue == "" {
				currentValue = mutantADN[r][c]
			}

			if currentValue == mutantADN[r][c+1] {
				coincidences++
			} else {
				currentValue = ""
			}

			if coincidences == 4 {
				countValid++
				coincidences = 0
			}
		}
	}

	return countValid
}

func vertical(mutantADN [][]string) int {
	countValid := 0
	for i := 0; i < len(mutantADN); i++ {
		coincidences := 1
		currentValue := ""

		for j := 0; j < len(mutantADN)-1; j++ {
			if currentValue == "" {
				currentValue = mutantADN[j][i]
			}

			if currentValue == mutantADN[j+1][i] {
				coincidences++
			} else {
				currentValue = ""
			}

			if coincidences == 4 {
				countValid++
				coincidences = 0
			}
		}
	}

	return countValid
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

	result := isMutant(data)

	if result {
		fmt.Println("Result:", "Es mutante")
	} else {
		fmt.Println("Result:", "No es mutante")
	}
}
