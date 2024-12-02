package processors

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Registry[1] = ProcessDay1
}

func computeLists(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")

	var arr1 []int
	var arr2 []int

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) >= 2 {
			num1, err1 := strconv.Atoi(fields[0])
			num2, err2 := strconv.Atoi(fields[1])

			if err1 != nil || err2 != nil {
				return nil, nil, errors.New("Cannot Convert to numbers")
			}
			arr1 = append(arr1, num1)
			arr2 = append(arr2, num2)
		}
	}

	return arr1, arr2, nil
}

func ProcessDay1(input string) (interface{}, interface{}, error) {
	result1, _ := ProcessDay1Part1(input)
	result2, _ := ProcessDay1Part2(input)

	return result1, result2, nil
}

func ProcessDay1Part2(input string) (interface{}, error) {
	var result int = 0

	arr1, arr2, err := computeLists(input)
	if err != nil {
		return nil, err
	}

	occurrences1 := make(map[int]int)
	occurrences2 := make(map[int]int)

	for _, num := range arr1 {
		occurrences1[num]++
	}

	for _, num := range arr2 {
		occurrences2[num]++
	}

	for num, frequency := range occurrences1 {
		result += num * frequency * occurrences2[num]
	}

	return result, nil
}

func ProcessDay1Part1(input string) (interface{}, error) {
	var result int = 0

	arr1, arr2, err := computeLists(input)
	if err != nil {
		return nil, err
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < len(arr1); i++ {
		diff := arr2[i] - arr1[i]
		if diff < 0 {
			diff = -diff
		}

		result += diff
	}

	return result, nil
}
