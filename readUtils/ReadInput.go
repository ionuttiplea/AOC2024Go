package readUtils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func requestInput(dayNo int) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", dayNo)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Printf("Error creating GET request: %v\n", err)
		return nil, errors.New("ERR")
	}

	sessionCookie := os.Getenv("SESSION_COOKIE")

	if sessionCookie == "" {
		log.Println("Session Cookie env variable not set")
		return nil, errors.New("ERR")
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionCookie,
		Path:  "/",
	}

	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Error making request: %v\n", err)
		return nil, errors.New("ERR")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, errors.New("ERR")
	}

	return body, nil
}

func writeToFile(fileName string, data []byte) error {
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		log.Printf("Cannot create directory %s : %v\n", dir, err)

	}

	err = os.WriteFile(fileName, data, 0644)

	if err != nil {
		log.Printf("Cannot write to file %s : %v\n", fileName, err)
		return err
	}

	return nil
}

func GetInput(dayNo int) (string, error) {
	inputFile := fmt.Sprintf("./inputs/Day%d.txt", dayNo)
	content, err := os.ReadFile(inputFile)

	if err != nil {
		log.Println("Input file not found, making request to AOC to retrieve it")

		input, err := requestInput(dayNo)
		if err != nil {
			return "", err
		}

		err = writeToFile(inputFile, input)
		if err != nil {
			return "", err
		}

		log.Println("Input file created, next time the input will be retrieved from this file")

		return string(input), nil
	}

	log.Printf("Retrieved input from : %s\n", inputFile)

	return string(content), nil
}
