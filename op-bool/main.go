package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 88

	var lulusAkhir bool = nilaiAkhir > 89
	var lulusAbsensi bool = absensi > 89
	fmt.Println(lulusAkhir)
	fmt.Println(lulusAbsensi)

	var lulus bool = lulusAkhir && lulusAbsensi
	fmt.Println(lulus)
}
