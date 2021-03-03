package main

import "fmt"

func main()  {
	fmt.Println( ">>> Conversi Integer")
	
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32, ">>> 32 Bit")
	fmt.Println(nilai64, ">>> 64 bit")
	fmt.Println(nilai8, ">>> 8 bit batasnya 127")
	
	fmt.Println( "\n>>> Conversi String")
	var name = "Abdul Rozak"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name, ">>> name")
	fmt.Println(eString, ">>> string(name[0])")
}