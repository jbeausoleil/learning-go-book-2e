package main

import "fmt"

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func main() {
	fmt.Println(MakePerson("Bob", "Smith", 30))
	fmt.Println(MakePersonPointer("Jim", "Salty", 30))
}

func MakePerson(firstname, lastname string, age int) Person {
	return Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}

func MakePersonPointer(firstname, lastname string, age int) *Person {
	return &Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}
