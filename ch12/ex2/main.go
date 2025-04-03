package main

import (
	"fmt"
	"math/rand"
)

func numGenerator(amt int) []int {
	nums := make([]int, amt)
	for i := 0; i < amt; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}

func ProcessData(nums []int) {
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		for _, v := range nums {
			ch <- v
		}
		close(ch)
	}()
	go func() {
		for _, v := range nums {
			ch2 <- v
		}
		close(ch2)
	}()
	vals := readFromChannels(ch, ch2)
	fmt.Println(vals)
}

func readFromChannels(ch, ch2 chan int) []int {
	var out []int
	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				count--
				break
			}
			out = append(out, v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			out = append(out, v)
		}
	}
	return out
}

func main() {
	nums := numGenerator(10)
	ProcessData(nums)
}
