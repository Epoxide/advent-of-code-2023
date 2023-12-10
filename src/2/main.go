package main

import (
	"advent-of-code-2023/2/game"
	"log"
	"os"
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

	for i, s := range games {
		if game.PossibleGame(game.MaxCubes(s), allowed) {
			sum += i + 1
		}
	}

	println("Part 1:", sum)
}
