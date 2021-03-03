package main

import "fmt"

func main() {
	// similar to var but will not get error if not used
	// const firstName string = "Abdul"
	// const lastName = "Rozak"
	// const value = 1000

	const (
		firstName string = "Abdul"
		lastName = "Rozak"
		value = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}