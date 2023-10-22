package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Contoh variabel dengan tipe data int
	var number int = 42
	// Menggunakan reflect.TypeOf() untuk mendapatkan tipe data
	typeOfNumber := reflect.TypeOf(number)

	var b int = 388
	//typeOfPloat := reflect.TypeOf(ploat)

	fmt.Printf("Tipe data dari 'number' adalah: %s\n", typeOfNumber)
	fmt.Printf("type data:%T\n", b)
	//fmt.Printf("Tipe data dari 'number' adalah: %s\n", typeOfPloat)
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)
}
