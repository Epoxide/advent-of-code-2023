package decode

import (
	"log"
	"regexp"
	"strconv"
)

func Decode(calibrationValue string) int {
  matchDigits := regexp.MustCompile(`\d{1,1}`)
	digits := matchDigits.FindAllString(calibrationValue, -1)
	digit := string(digits[0]) + string(digits[len(digits)-1])
	output, err := strconv.Atoi(digit)

	if err != nil {
		log.Fatal(err)
	}

	return output
}
