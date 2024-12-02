package processors

import (
	"log"
	"strconv"
	"strings"
)

func init() {
	Registry[2] = ProcessDay2
}

func ProcessDay2(input string) (interface{}, interface{}, error) {
	result1, _ := ProcessDay2Part1(input)
	result2, _ := ProcessDay2Part2(input)

	return result1, result2, nil
}

func ProcessDay2Part2(input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	safeCount := 0

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		numbers := make([]int, 0)

		for _, field := range fields {
			num, err := strconv.Atoi(field)

			if err != nil {
				log.Println("Cannot convert to integer")
				return nil, err
			}
			numbers = append(numbers, num)
		}

		safe, firstBadIndex := isSafe(numbers)

		if safe {
			safeCount++
			continue
		}

		newNumbers1 := make([]int, 0)
		newNumbers1 = append(newNumbers1, numbers[:firstBadIndex]...)
		newNumbers1 = append(newNumbers1, numbers[firstBadIndex+1:]...)

		safe, _ = isSafe(newNumbers1)

		if safe {
			safeCount++
			continue
		}

		newNumbers2 := make([]int, 0)
		newNumbers2 = append(newNumbers2, numbers[:firstBadIndex-1]...)
		newNumbers2 = append(newNumbers2, numbers[firstBadIndex:]...)

		safe, _ = isSafe(newNumbers2)
		if safe {
			safeCount++
			continue
		}

		newNumbers3 := make([]int, 0)
		newNumbers3 = append(newNumbers3, numbers[1:]...)

		safe, _ = isSafe(newNumbers3)
		if safe {
			safeCount++
			continue
		}
	}

	return safeCount, nil
}

func ProcessDay2Part1(input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	safeCount := 0

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		numbers := make([]int, 0)

		for _, field := range fields {
			num, err := strconv.Atoi(field)

			if err != nil {
				log.Println("Cannot convert to integer")
				return nil, err
			}
			numbers = append(numbers, num)
		}

		safe, _ := isSafe(numbers)

		if safe {
			safeCount++
			continue
		}
	}

	return safeCount, nil
}

func isSafe(numbers []int) (bool, int) {
	monotony := ""

	if !isSafeDiff(numbers[0], numbers[1]) {
		return false, 1
	}

	if len(numbers) > 1 && numbers[0] > numbers[1] {
		monotony = "desc"
	}

	if len(numbers) > 1 && numbers[0] < numbers[1] {
		monotony = "asc"
	}

	if len(numbers) > 1 && numbers[0] == numbers[1] {
		return false, 1
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] && monotony == "desc" {
			return false, i
		}

		if numbers[i] < numbers[i-1] && monotony == "asc" {
			return false, i
		}

		if !isSafeDiff(numbers[i], numbers[i-1]) {
			return false, i
		}
	}

	return true, -1
}

func isSafeDiff(a int, b int) bool {
	diff := a - b
	return diff != 0 && diff <= 3 && diff >= -3
}
