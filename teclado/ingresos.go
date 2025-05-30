package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error()) // El panic se utiliza para abortar la operacion

		}
	}

	fmt.Println("Ingrese número 2")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error()) // El panic se utiliza para abortar la operacion

		}
	}

	fmt.Println("Ingrese Leyenda")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
