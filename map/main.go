package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "ata",
		"addres": "sumber",
	}
	fmt.Println(person)

	person["name"] = "tuti" //mengganti nilai data map
	var per = person["name"]
	fmt.Println(per)

	fmt.Println(len(person))      //jumlah map
	fmt.Println(person["addres"]) //melihat nilai map
	//membuat make map baru
	var book = make(map[string]string)
	book["author"] = "jamal"
	book["title"] = "golang"

	fmt.Println(book)
}
