package main

import (
	"fmt"
	"latihan/calculation"
)

func main() {
	hasil_tambah := calculation.Add(10, 2)
	hasil_kali := calculation.Time(10, 2)
	fmt.Println(hasil_tambah)
	fmt.Println(hasil_kali)
}
