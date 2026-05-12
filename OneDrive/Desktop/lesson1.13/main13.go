package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль")
	}
	return a / b, nil
}

func main() {
	res, err := divide(10.5, 2.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)
	}
}
