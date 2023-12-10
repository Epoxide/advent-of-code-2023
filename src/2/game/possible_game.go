package game

import "strconv"

func PossibleGame(game map[string]string, allowed map[string]string) bool  {
	for color, amount := range game {
		amountInt, _ := strconv.Atoi(amount)
		allowedInt, _ := strconv.Atoi(allowed[color])

		if allowedInt < amountInt {
			return false
		}
	}

	return true
}
