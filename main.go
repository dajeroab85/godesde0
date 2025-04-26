package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	//fmt.Println("Hola mundo")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)
}
