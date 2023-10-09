package lenguaje

import (
	"fmt"
)

type Environment struct {
	Anterior interface{}
	Tabla    map[string]Symbol
	Id       string
	Size     map[string]int
}

func NewEnvironment(ant interface{}, id string) Environment {
	fmt.Println("Creando nuevo entorno")
	env := Environment{
		Anterior: ant,
		Tabla:    make(map[string]Symbol),
		Id:       id,
		Size:     make(map[string]int),
	}
	env.Size["size"] = 0
	return env
}

func (env Environment) SaveVariable(id string, tipo TipoExpresion, constante bool, linea int, columna int) Symbol {
	if variable, ok := env.Tabla[id]; ok {
		fmt.Println("La variable "+id+" ya existe ", variable)
		return Symbol{Lin: 0, Col: 0, Tipo: NIL, Constante: false, Posicion: 0}
	}
	env.Tabla[id] = Symbol{Lin: linea, Col: columna, Tipo: tipo, Constante:constante, Posicion: env.Size["size"]}
	env.Size["size"] = env.Size["size"] + 1
	return env.Tabla[id]
}

func (env Environment) GetVariable(id string) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NIL, Posicion: 0}
}

func (env Environment) SetVariable(id string, value Symbol) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			tmpEnv.Tabla[id] = value
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NIL, Posicion: 0}
}