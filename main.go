package main

import (
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input/2022011")
	if err != nil {
		log.Fatalf("can't open input: %v", err)
	}

	result, err := AOC2022011(string(content))
	if err != nil {
		log.Fatalf("can't parse input: %v", err)
	}

	log.Printf("AOC2022 01 1: %d", result)

	result, err = AOC2022012(string(content))
	if err != nil {
		log.Fatalf("can't parse input: %v", err)
	}

	log.Printf("AOC2022 01 2: %d", result)
}
