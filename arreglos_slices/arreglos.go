package arreglos_slices

import "fmt"

// Los Slices no tienen una cantidad defifinida de elementos mientras que los arreglos se deben definir la dimensión

var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"Daniel", "Pablo", "Juan"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
