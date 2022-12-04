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

	content, err = os.ReadFile("input/2022021")
	if err != nil {
		log.Fatalf("can't open input: %v", err)
	}

	result, err = AOC2022021(string(content))
	if err != nil {
		log.Fatalf("can't parse input: %v", err)
	}

	log.Printf("AOC2022 02 1: %d", result)

	result, err = AOC2022022(string(content))
	if err != nil {
		log.Fatalf("can't parse input: %v", err)
	}

	log.Printf("AOC2022 02 2: %d", result)

	content, err = os.ReadFile("input/2022031")
	if err != nil {
		log.Fatalf("can't open input: %v", err)
	}

	result, err = AOC2022031(string(content))
	if err != nil {
		log.Fatalf("can't parse input: %v", err)
	}

	log.Printf("AOC2022 03 1: %d", result)

	result, err = AOC2022032(string(content))
	if err != nil {
		log.Fatalf("can't parse input: %v", err)
	}

	log.Printf("AOC2022 03 2: %d", result)
}
