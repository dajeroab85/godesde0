package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabladeMultiplicar() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin) // Inicializa el obejo scanner para leer desde el teclado
	for {
		fmt.Println("Ingrese un  número")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue

			} else {
				break
			}
		}
	}

	for i := 1; i < 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	return texto
}
