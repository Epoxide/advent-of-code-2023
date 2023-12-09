package letterdigits

import "strings"

func DigitLookup(input string) string {
	var output string

	if len(input) == 1 {
		return input
	}

	numbers := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
		"zero": "0",
	}

	for key, value := range numbers {
		if strings.Contains(input, key) {
			output = value
		}
	}

	return output
}
