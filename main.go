package main

import (
	"github.com/godesde0/files"
)

func main() {
	//fmt.Println("Hola mundo")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)*/

	// Los if permiten hecer asignaciones dentro de su instruccion. Sintaxis if asignacion; pregunta del if

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es:", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	// Al igual que el if el Swicht me permite hacer la asignacion y posterior la pregunta. Sintaxis switch asignacion; pregunta
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/

	/*

		numero, texto := ejercicios.ConviertoEntero("500")

		fmt.Println(numero)
		fmt.Println(texto)*/

	//teclado.IngresoNumero()

	// Esto sería equivalente a usar un while infinito

	//iteraciones.Iterar()

	//fmt.Println(ejercicios.TabladeMultiplicar())

	//files.GrabaTabla()

	//files.SumaTabla()

	files.LeoArchivo()
}
