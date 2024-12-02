package processors

import "fmt"

func init() {
	Registry[0] = NotYetImplementedProcess
}

func NotYetImplementedProcess(input string) (interface{}, interface{}, error) {
	return nil, nil, fmt.Errorf("Not Created")
}
