package simple_parser

type Symbol struct {
	VariableName string
	Type         string
}

func NewSymbol(name, var_type string) *Symbol {
	return &Symbol{
		VariableName: name,
		Type:         var_type,
	}
}
