package main

import "fmt"

func main() {
	//diketahui panjang array
	days := [7]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days[0])
	fmt.Println(len(days))

	//days[1] = "libur"

	//tidak diketahui panjang array
	day := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(day)
}
