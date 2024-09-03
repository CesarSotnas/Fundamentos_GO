package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// mais de um retorno na mesma função
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	fmt.Println("Teste de função")

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println(text)
		return text
	}

	resultado := f("Texto da função 2")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)

	_, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao)

}
