package processors

import "fmt"

func init() {
	Registry[0] = NotYetImplementedProcess
}

func NotYetImplementedProcess(input string) (interface{}, error) {
	return nil, fmt.Errorf("Not Created")
}
