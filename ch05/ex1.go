package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	expressions := [][]string{
		{"1", "+", "2"},
		{"1", "-", "2"},
		{"1", "*", "2"},
		{"1", "/", "0"},
	}

	for _, v := range expressions {
		if len(v) != 3 {
			fmt.Println("Invalid expression")
			continue
		}
		opFunc, ok := operationsMap[v[1]]
		if !ok {
			fmt.Println("Invalid operation", v[1])
			continue
		}
		p1, err := strconv.Atoi(v[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(v[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p1, v[1], p2, "=", result)
	}
}

func add(x, y int) (int, error) {
	return x + y, nil
}

func sub(x, y int) (int, error) {
	return x - y, nil
}

func mul(x, y int) (int, error) {
	return x * y, nil
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

var operationsMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
