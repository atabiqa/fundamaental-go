package main

import "fmt"

func main() {
	//variabel konstanta
	const firstName string = "atabiq"
	const lastName = "as'ad"
	const value = 1000
	const bird = 70 //tidak dipanggil tidak error

	//multipel variabel
	const (
		panggilan = "atabiq"
		umur      = 20
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
