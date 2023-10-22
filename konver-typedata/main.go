package main

import "fmt"

func main() {
	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	//jika tidak pas type datanya saat konversi maka akan berulang

	//ex konversi 2
	name := "ata"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println(name)
	fmt.Println(eString)
}
