package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
