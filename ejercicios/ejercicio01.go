package ejercicios

import "strconv"

func ConviertoEntero(valorstring string) (int, string) {

	valorint, err := strconv.Atoi(valorstring)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if valorint > 100 {
		return valorint, "Es mayor a 100"
	} else {
		return valorint, "Es menor a 100"
	}

}
