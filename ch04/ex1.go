package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var intSlice []int
	for i := 0; i < 10; i++ {
		intSlice = append(intSlice, rand.Intn(100))
	}
	fmt.Println(intSlice)
}
