package main

import "fmt"

func main() {
	sTest := []string{"a", "b", "c"}
	fmt.Println(sTest)
	UpdateSlice(sTest, "d")
	GrowSlice(sTest, "e")
	fmt.Println(sTest)
}

func UpdateSlice(ss []string, s string) {
	ssLength := len(ss)
	ss[ssLength-1] = s
	fmt.Println(ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss)
}
