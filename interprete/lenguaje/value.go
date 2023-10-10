package lenguaje

import "fmt"

type Value struct {
	Value      string
	IsTemp     bool
	Type       TipoExpresion
	TrueLabel  []interface{}
	FalseLabel []interface{}
	OutLabel   []interface{}
	IntValue   int
	StringValue string
}

func NewValue(Val string, tmp bool, typ TipoExpresion, strVal string) Value {
	result := Value{
		Value:    Val,
		IsTemp:   tmp,
		Type:     typ,
		IntValue: 0,
		StringValue: strVal,
	}
	return result
}

func getDataValue(v Value) {
	fmt.Println("--- Value ---")
	fmt.Println("Value: ", v.Value)
	fmt.Println("IsTemp: ", v.IsTemp)
	fmt.Println("Type: ", v.Type)
	fmt.Println("TrueLabel: ", v.TrueLabel)
	fmt.Println("FalseLabel: ", v.FalseLabel)
	fmt.Println("OutLabel: ", v.OutLabel)
	fmt.Println("IntValue: ", v.IntValue)
	fmt.Println("--------------")
}