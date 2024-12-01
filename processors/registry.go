package processors

var Registry = make(map[int]func(string) (interface{}, error))

func GetProcessor(day int) (func(string) (interface{}, error), error) {
	processFunc, exists := Registry[day]
	if !exists {
		processFunc = Registry[0]
	}

	return processFunc, nil
}
