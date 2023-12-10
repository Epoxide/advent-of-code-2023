package game

import (
	"strconv"
	"strings"
)

func MaxCubes(input string) map[string]string {
	sets := strings.SplitAfter(input, ": ")[1]
	splitSets := strings.Split(sets, "; ")

	maxCubes := map[string]string{
		"red":    "0",
		"green":  "0",
		"blue":   "0",
	}

	for _, s := range splitSets {
		colorSplit := strings.Split(s, ", ")

		for _, c := range colorSplit {
			amountSplit := strings.Split(c, " ")
			amount := amountSplit[0]
			color := amountSplit[1]
			amountInt, _ := strconv.Atoi(amount)
			currentAmountInt, _ := strconv.Atoi(maxCubes[color])

			if currentAmountInt < amountInt {
				maxCubes[color] = amount
			}			
		}
	}

	return maxCubes
}
