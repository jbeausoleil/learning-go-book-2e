package main

import "fmt"

// Write a generic function that doubles the value
// of any integer or float that's passed into it.
// Define any needed generic interfaces.

type ValidType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Double[T ValidType](x T) T {
	return x * 2
}

func main() {
	fmt.Println(Double(3))
	fmt.Println(Double(3.14))
}
