package funciones

import "fmt"

func Exponecia(valor int) {
	if valor > 1000000000 {
		return
	}

	fmt.Println(valor)
	Exponecia(valor * 2)
}
