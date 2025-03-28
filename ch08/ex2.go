package main

import (
	"fmt"
	"strconv"
)

// Define a generic interface called Printable that matches a type that implements
// fmt.Stringer and has an underlying type of int or float64.  Define types that meet this interface.
// Write a function that takes in a Printable and prints its value to the screen using fmt.Println

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return strconv.FormatFloat(float64(pf), 'f', -1, 64)
}

func Print[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	a := PrintInt(10)
	b := PrintFloat(10.2)
	Print(a)
	Print(b)
}
