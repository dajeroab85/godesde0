package modelos

// Herencia
// El modelo Mujer hereda la estructura de Hombre
type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Comer() {
	m.comiendo = true
}

func (m *Mujer) Pensa() {
	m.pensando = true
}

func (m *Mujer) Sexo() string {
	return "Mujer"
}
