package inter

type NodeInterface interface {
	Errors(s string) error
	NewLabel() uint32
	EmitLabel(i uint32)
	Emit(code string)
}

type ExprInterface interface {
	NodeInterface
	Gen() ExprInterface
	Reduce() ExprInterface
	Type() *Type
	ToString() string
}
