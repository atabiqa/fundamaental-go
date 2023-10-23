package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5} // slice adalah ketika dalam kurung kotak tidak ada isi
	fmt.Println(slice)

	//slice dengan array
	moon := [...]string{
		"januari",
		"feb",
		"mar",
		"april",
		"mei",
		"juni",
		"juli",
		"agus",
	}

	fmt.Println(moon) // menampilkan array

	var moon1 = moon[2:5] // memanggil slice
	var moon2 = moon[:]
	var moon3 = moon[:5]

	fmt.Println(moon1)
	fmt.Println(len(moon1))
	fmt.Println(cap(moon1))
	moon1[0] = "liburan" // merubah slice
	fmt.Println(moon1)   //menampilkan perubahan slice

	fmt.Println(moon2)
	fmt.Println(moon3)

	newslice := make([]string, 2, 5)
	newslice[0] = "oki"
	newslice[1] = "toi"
	//newslice[2] = "asa"

	fmt.Println(newslice)

}
