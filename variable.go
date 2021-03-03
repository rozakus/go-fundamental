package main

import "fmt"

func main() {
	// declare dataype
	fmt.Println("Declare Var Using Datatypes")
	fmt.Println(">>> var name string")
	
	var name string
	
	name = "Abdul Rozak 1"
	fmt.Println(name)
	
	name = "Abdul Rozak 2"
	fmt.Println(name)
	
	// declare dataype without define it, auto
	fmt.Println("\nDeclare Var WITHOUT Using Datatypes")
	fmt.Println(">>> var friendName = Abdul Rozak 3")
	
	var friendName = "Abdul Rozak 3"
	fmt.Println(friendName)
	
	// int32/int64 as default due to OS or declare the type
	var age = 25
	fmt.Println(age)
	
	// declare without var
	fmt.Println("\nDeclare Var WITHOUT Using Var & Datatypes")
	fmt.Println(">>> country := Indonesia")
	country := "Indonesia"
	fmt.Println(country)

	// Declare multiple var
	fmt.Println("\nDeclare Multiple Var")
	fmt.Println(">>> var ( \nfirstName = Abdul\nlastName = Rozak\n)")

	var (
		firstName = "Abdul"
		lastName = "Rozak"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}