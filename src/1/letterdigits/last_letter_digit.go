package letterdigits

import "strings"

func LastLetterDigit(input string) string {
	numbers := [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	var lastOccurence string
	lastOccurenceIndex := 0

	for _, number := range numbers {
		if strings.LastIndex(input, number) > lastOccurenceIndex {
			lastOccurenceIndex = strings.LastIndex(input, number)
			lastOccurence = number 
		}
	}

	return lastOccurence
}
