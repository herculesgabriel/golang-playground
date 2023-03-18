package utils

import "errors"

func Divide(a int, b int) (int, error) {
	if a == 0 {
		privateLogger("There is an error!")
		return 0, errors.New("first value cannot be zero")
	}
	return a / b, nil
}
