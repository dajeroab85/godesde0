package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {

	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   // Slide creado desde un vector, desde la posición 3
	porcion2 := arreglo[:5]  // slide creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // slide creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capadidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capadidad %d", len(nums), cap(nums))
}
