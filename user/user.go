package user

import (
	"fmt"
	"time"

	"github.com/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)

	u.AddUser(10, "Daniel", time.Now(), true)

	fmt.Println(u)
}
