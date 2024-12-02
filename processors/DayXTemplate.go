package processors

func init() {
	Registry[0] = NotYetImplementedProcess
}

func ProcessDayX(input string) (interface{}, interface{}, error) {
	result1, _ := ProcessDayXPart1(input)
	result2, _ := ProcessDayXPart2(input)

	return result1, result2, nil
}

func ProcessDayXPart1(input string) (interface{}, error) {
	return nil, nil
}

func ProcessDayXPart2(input string) (interface{}, error) {
	return nil, nil
}
