package game

import (
	"testing"
)

func TestMaxCubes(t *testing.T) {
  result := MaxCubes("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	maxCubes := map[string]string{
		"red":    "4",
		"green":  "2",
		"blue":   "6",
	}

	if result["red"] != maxCubes["red"] || result["green"] != maxCubes["green"] || result["blue"] != maxCubes["blue"] {
    t.Error("Incorrect max cubes", result)
  }

	result = MaxCubes("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")

	maxCubes = map[string]string{
		"red":    "20",
		"green":  "13",
		"blue":   "6",
	}

	if result["red"] != maxCubes["red"] || result["green"] != maxCubes["green"] || result["blue"] != maxCubes["blue"] {
    t.Error("Incorrect max cubes", result)
  }
}
