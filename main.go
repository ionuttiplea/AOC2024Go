package main

import (
	"adventOfCode2024/processors"
	"adventOfCode2024/readUtils"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Cannot read .env file, make sure you have .env in the root of the project")

		return
	}

	dayNo, err := strconv.Atoi(os.Getenv("DAY_NUMBER"))

	if err != nil {
		log.Println("DAY_NUMBER is not a valid integer")
		return
	}

	if dayNo > 31 || dayNo < 1 {
		log.Println("DAY_NUMBER can only be a number between 1 and 31")
		return
	}
	log.Println("Env Variables loaded, proceeding ...")

	input, err := readUtils.GetInput(dayNo)

	if err != nil {
		return
	}

	processFunc, err := processors.GetProcessor(dayNo)

	if err != nil {
		return
	}

	log.Printf("Starting processing input for Day %d ...\n", dayNo)
	start := time.Now()

	result1, result2, err := processFunc(input)

	if err != nil {
		log.Println("An error has occured while processing:", err)
		return
	}

	elapsed := time.Since(start)
	log.Print("Result part 1 : ", result1)
	log.Print("Result part 2 : ", result2)

	log.Printf("The algorithm used to solve the problem took %s\n", elapsed)

	log.Println("Done.")
}
