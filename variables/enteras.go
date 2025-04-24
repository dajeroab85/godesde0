package variables

import "fmt"

// Para que una función sea declarada como Pública y accedida desde afuera debe ser declarada en Mayuscula
func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun= ", intcomun)
	fmt.Println("intde32= ", intde32)
	fmt.Println("intde64= ", intde64)

}
