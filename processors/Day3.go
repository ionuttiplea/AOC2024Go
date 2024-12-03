package processors

import (
	"regexp"
	"strconv"
)

func init() {
	Registry[3] = ProcessDay3
}

func ProcessDay3(input string) (interface{}, interface{}, error) {
	result1, _ := ProcessDay3Part1(input)
	result2, _ := ProcessDay3Part2(input)

	return result1, result2, nil
}

func ProcessDay3Part1(input string) (interface{}, error) {
	return computeMulsOnString(input), nil
}

func ProcessDay3Part2(input string) (interface{}, error) {
	splitByDosAndDontsRegex := regexp.MustCompile(`do\(\)|don't\(\)`)
	result := int64(0)

	parts := splitByDosAndDontsRegex.Split(input, -1)
	doesAndDonts := splitByDosAndDontsRegex.FindAllString(input, -1)

	for i, part := range parts {
		if i == 0 {
			result += computeMulsOnString(part)
			continue
		}

		if doesAndDonts[i-1] == "do()" {
			result += computeMulsOnString(part)
		}
	}

	return result, nil
}

func computeMulsOnString(input string) int64 {
	extractMulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := extractMulRegex.FindAllStringSubmatch(input, -1)

	result := int64(0)

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		result += int64(num1 * num2)
	}

	return result

}
