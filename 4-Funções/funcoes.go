package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	divisao := n1 / n2
	multiplicacao := n1 * n2
	return soma, subtracao, divisao, multiplicacao

}

func main() {
	// cal := calculosMatematicos(12, 10)
	// fmt.Println(cal)

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função1 ")
	fmt.Println(resultado)

	resulSoma, resulSub, resulDiv, resulMult := calculosMatematicos(10, 90)
	fmt.Println(resulSoma, resulSub, resulDiv, resulMult)

	//usar o _ para dizer que quer para ignorar um parametro.
}
