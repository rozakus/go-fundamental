package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	fmt.Println("ujian >= 80 : ", lulusUjian)
	fmt.Println("absensi >= 80 : ",lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println("lulusAbsensi && lulusUjian : ",lulus)

	fmt.Println("ujian >= 80 && absensi >= 80 : ",ujian >= 80 && absensi >= 80)
}