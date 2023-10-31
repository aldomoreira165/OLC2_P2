package lenguaje

type Error struct {
	Description string
	Linea       int
	Columna     int
}

type TablaErrores struct {
	Tabla map[int]Error
}

func NewTablaErrores() *TablaErrores {
	return &TablaErrores{
		Tabla: make(map[int]Error),
	}
}

func (te *TablaErrores) InsertarError(description string, linea int, columna int) {
	te.Tabla[linea] = Error{description, linea, columna}
}