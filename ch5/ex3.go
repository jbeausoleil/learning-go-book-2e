package main

import "fmt"

func main() {
	helloPrefix := prefixer("Hello ")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return prefix + " " + body
	}
}
