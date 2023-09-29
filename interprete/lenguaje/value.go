package lenguaje

import (
	arrayList "interprete/lenguaje/arraylist"
)

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
	TrueLabel    *arrayList.List
	FalseLabel   *arrayList.List
	OutLabel     *arrayList.List
}

func NewValue(Val string, tmp bool, typ TipoExpresion) Value {
	result := Value{
		Value:        Val,
		IsTemp:       tmp,
		Type:         typ,
		TrueLabel:    arrayList.New(),
		FalseLabel:   arrayList.New(),
		OutLabel:     arrayList.New(),
	}
	return result
}