package main

import (
	"advent22/day5"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// day1.Day1()
	// day2.Day2()
	// day3.Day3Part2()
	// day4.Day4()
	day5.Day5()
}
