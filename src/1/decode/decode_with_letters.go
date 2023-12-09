package decode

import (
	"advent-of-code-2023/1/letterdigits"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func DecodeWithLetters(calibrationValue string) int {
  matchDigitsIncludingDigitsSpelled := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine|zero`)
	digits := matchDigitsIncludingDigitsSpelled.FindAllString(calibrationValue, -1)
	lastDigit := letterdigits.LastLetterDigit(calibrationValue)

	if lastDigit == "" || strings.LastIndex(calibrationValue, lastDigit) < strings.LastIndex(calibrationValue, digits[len(digits) - 1]) {
		lastDigit = digits[len(digits) - 1]
	}

	digit := letterdigits.DigitLookup(digits[0]) + letterdigits.DigitLookup(lastDigit)
	output, err := strconv.Atoi(digit)

	if err != nil {
		log.Fatal(err)
	}

	return output
}
