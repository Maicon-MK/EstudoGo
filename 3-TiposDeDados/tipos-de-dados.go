package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000
	fmt.Println(numero)

	var numero2 uint32 = 15647
	fmt.Println(numero2)

	//alias
	//INT32
	var numero3 rune = 12354
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 1526.25
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1526.25
	fmt.Println(numeroReal2)

	numeroReal3 := 1587.55
	fmt.Println(numeroReal3)
	//O char vai gerar um numero de acordo com o caracter da tabela ASCII. No caso de B ele retorna 66, no b minusculo 98
	//Mas e preciso ultizar as aspas simples, aspas duplas somente para Str
	char := 'B'
	fmt.Println(char)

	//BOOOLEANOL

	var booleano1 bool //FALSE
	fmt.Println(booleano1)

	var booleano2 bool = true //TRUE
	fmt.Println(booleano2)

	//ERRO

	var erro error
	fmt.Println(erro) //nil (null)

	var erro2 error = errors.New("Erro aqui")
	fmt.Println(erro2)
}
