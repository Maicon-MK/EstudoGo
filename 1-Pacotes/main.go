package main

import (
	"fmt"
	"modulo/1-Pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Arquivo main")
	auxiliar.Escrever()
	checkmail.ValidateFormat("joaomichaelfd@gmail.com")

	error := checkmail.ValidateFormat("joaomichaelfdgmail.com")

	fmt.Print(error, "O email esta correto")
}
