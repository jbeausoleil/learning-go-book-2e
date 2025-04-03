package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
1. Create a function that launches three goroutines that communicate using a channel.
The first two goroutines each write 10 numbers to the channel.
The third goroutine reads all the numbers from the channel and prints them out.
The function should exit when all values have been printed out.
Make sure that none of the goroutines leak.
You can create additional goroutines, if needed.
*/

func numGenerator(amt int) []int {
	nums := make([]int, amt)
	for i := 0; i < amt; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}

func ProcessData(nums []int, wgWriteNum, wgReadNum int) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(wgWriteNum)
	for i := 0; i < wgWriteNum; i++ {
		go func() {
			defer wg.Done()
			for _, v := range nums {
				ch <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(wgReadNum)
	for i := 0; i < wgReadNum; i++ {
		go func() {
			defer wg2.Done()
			for v := range ch {
				fmt.Println(v)
			}
		}()
	}
	wg2.Wait()
}

func main() {
	nums := numGenerator(10)
	fmt.Println(nums)
	ProcessData(nums, 2, 1)
}
