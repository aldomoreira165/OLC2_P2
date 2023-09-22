package lenguaje

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
	ArrType      TipoExpresion
	ReturnFlag   bool
	BreakFlag    bool
	ContinueFlag bool
	IntValue     int
}

func NewValue(Val string, tmp bool, typ TipoExpresion, ret, br, cont bool) Value {
	result := Value{
		Value:        Val,
		IsTemp:       tmp,
		Type:         typ,
		ArrType:      NULL,
		ReturnFlag:   ret,
		BreakFlag:    br,
		ContinueFlag: cont,
		IntValue:     0,
	}
	return result
}