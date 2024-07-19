package main

import "fmt"

func main() {
	var variavel1 string = "hello Word"
	variavel2 := "ai papai."
	fmt.Println(variavel1, variavel2)
	//explicito
	var nome string = "Joao michael"
	var idade int = 26
	var funcao string = "Dev Full Stack"

	fmt.Println("Meu nome e", nome, "eu tenho", idade, "sou", funcao)

	//implicito
	nome2 := "joao mcihael"
	idade2 := 26
	funcao2 := "dev Full Stack"

	fmt.Println("Meu nome e", nome2, "tenho", idade2, "sou", funcao2)

	var (
		carro1 string = "UNO"
		carro2 string = "Celta"
	)
	fmt.Println(carro1, carro2)

	nome3, idade3 := "Michael", 16

	fmt.Println(nome3, idade3)

	//Todas formas de declaração pode ser usada com const
}
