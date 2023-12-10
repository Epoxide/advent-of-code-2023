package game

import (
	"testing"
)

func TestPossibleGame(t *testing.T) {
  game := MaxCubes("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	allowed := map[string]string{
		"red":    "12",
		"green":  "13",
		"blue":   "14",
	}

	result := PossibleGame(game, allowed)

	if result != true {
    t.Error("Game should be possible", result)
  }

	game = MaxCubes("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")

	allowed = map[string]string{
		"red":    "12",
		"green":  "13",
		"blue":   "14",
	}

	result = PossibleGame(game, allowed)

	if result != false {
    t.Error("Game should not be possible", result)
  }
}
