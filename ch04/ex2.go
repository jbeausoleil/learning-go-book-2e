package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var intSlice []int
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		intSlice = append(intSlice, num)
		switch {
		case num%2 == 0:
			fmt.Println(num, ":", "Divisible by 2")
		case num%3 == 0:
			fmt.Println(num, ":", "Divisible by 3")
		case num%2 == 0 && num%3 == 0:
			fmt.Println(num, ":", "Divisible by 2 and 3")
		default:
			fmt.Println(num, ":", "Never mind")
		}
	}
}
