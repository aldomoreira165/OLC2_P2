package lenguaje

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
}

func NewValue(Val string, tmp bool, typ TipoExpresion) Value {
	result := Value{
		Value:        Val,
		IsTemp:       tmp,
		Type:         typ,
	}
	return result
}