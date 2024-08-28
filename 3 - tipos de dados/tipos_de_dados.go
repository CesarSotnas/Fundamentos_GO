package main

import (
	"errors"
	"fmt"
)

func main() {
	// números inteiros int, int8, int16, int32, int64
	// uint = unsigned int

	var num int16 = 100
	var num2 uint = 10000

	fmt.Println(num)
	fmt.Println(num2)

	// alias
	// int32 = rune
	// int8 = byte

	var num3 rune = 123456
	var num4 byte = 123

	fmt.Println(num3)
	fmt.Println(num4)

	// floats - float32, float 64

	var numeroReal float32 = 132.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 44478941132.45
	fmt.Println(numeroReal2)

	// strings

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	// booleano

	var booleano bool = true
	fmt.Println(booleano)

	// erros

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
