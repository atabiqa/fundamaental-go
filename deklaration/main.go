package main

import "fmt"

func main() {
	type NoKtp string
	type merried bool

	var ktpAta NoKtp = "111111"
	var merriedStatus merried = true
	fmt.Println(ktpAta)
	fmt.Println(merriedStatus)
}
