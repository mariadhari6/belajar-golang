package main

import "fmt"

/*
	Tipe data  => int, float32, string, bool
*/
func main() {
	var nama string = "Mar'i Adhari" //static data type
	fmt.Println("Nama saya", nama)
	usia := 21 //dynamic data type {tipe data nya ditentikan oleh compiler berdasarkan nilai yang diberikan}
	fmt.Println("Saya berusia", usia, "Tahun")
}
