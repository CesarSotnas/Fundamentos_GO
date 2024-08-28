package main

import "fmt"

func main() {

	// inferência de tipo

	var variavel string = "variável 1"
	variavel2 := "variável 2"

	fmt.Println(variavel)
	fmt.Println(variavel2)

	variavel3, variavel4 := "variável 3", "variável 4"

	fmt.Println(variavel3)
	fmt.Println(variavel4)
}
