package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var usuario1 usuario
	usuario1.nome = "Dorival"
	usuario1.idade = 29
	fmt.Println(usuario1)

	enderecoExemplo := endereco{"rua dos Andradas", 15}

	usuario2 := usuario{"Percival", 68, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Rúbens"}
	fmt.Println(usuario3)
}
