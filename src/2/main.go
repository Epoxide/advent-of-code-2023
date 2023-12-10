package main

import (
	"advent-of-code-2023/2/game"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	
	games := strings.Split(string(content), "\n")
	
	allowed := map[string]string{
		"red":    "12",
		"green":  "13",
		"blue":   "14",
	}

	sum := 0
	sum2 := 0

	for i, s := range games {
		maxCubes := game.MaxCubes(s)

		if game.PossibleGame(maxCubes, allowed) {
			sum += i + 1
		}

		redAmount, _ := strconv.Atoi(maxCubes["red"])
		greenAmount, _ := strconv.Atoi(maxCubes["green"])
		blueAmount, _ := strconv.Atoi(maxCubes["blue"])
		product := redAmount * greenAmount * blueAmount
		sum2 += product
	}

	println("Part 1:", sum)
	println("Part 2:", sum2)
}
