package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

/*
Exercise 1. Create a sentinel error to represent an invalid ID. In +main()+, use +errors.Is+ to check for the
sentinel error, and print a message when it is found.
*/

/*
Create ErrInvalidID. Sentinel errors are package-level variables of type error that represent an error state.
Their names should start with Err.
*/

/*
errors.As:
It performs a type assertion through the error chain and assigns the error value to the provided variable
if it matches the type. This allows you to extract and work with any extra data contained in the custom error.

errors.Is:
It compares error values by unwrapping and checking if any error equals the target.
You don’t get access to any extra information the error might carry—just a boolean result.
*/

var ErrorInvalidId = errors.New("invalid ID")

type EmptyFieldError struct {
	Field string
}

func (fe EmptyFieldError) Error() string {
	return fe.Field
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		var fieldError EmptyFieldError
		if err != nil {
			if errors.Is(err, ErrorInvalidId) {
				fmt.Printf("record %d: %+v error: invalid ID: %s\n", count, emp, emp.ID)
			} else if errors.As(err, &fieldError) {
				fmt.Printf("record %d: %+v error: %s is empty\n", count, emp, fieldError.Field)
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			}
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return EmptyFieldError{Field: "ID"}
	}
	if !validID.MatchString(e.ID) {
		return ErrorInvalidId
	}
	if len(e.FirstName) == 0 {
		return EmptyFieldError{Field: "FirstName"}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldError{Field: "LastName"}
	}
	if len(e.Title) == 0 {
		return EmptyFieldError{Field: "Title"}
	}
	return nil
}
