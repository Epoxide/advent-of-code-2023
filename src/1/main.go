package main

import (
	"advent-of-code-2023/1/decode"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	
	calibrations := strings.Split(string(content), "\n")
	sum := 0

	for _, s := range calibrations {
		sum += decode.Decode(s)
	}

	println(sum)
}
