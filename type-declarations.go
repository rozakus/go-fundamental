package main

import "fmt"

func main() {
	type NOKTP string
	type MARRIED bool

	var noKtpRozak NOKTP = "1234567890"
	var marriedStatus MARRIED = false

	fmt.Println(noKtpRozak)
	fmt.Println(marriedStatus)
}