package simple_parser

type NodeInterface interface {
	AddChild(child NodeInterface)
	GetChildren() []NodeInterface
	Attribute() string
}

type SyntaxNode struct {
	T        string //节点属性
	children []NodeInterface
}

func NewSyntaxNode() *SyntaxNode {
	return &SyntaxNode{
		T: "",
	}
}

func (s *SyntaxNode) AddChild(node NodeInterface) {
	s.children = append(s.children, node)
}

func (s *SyntaxNode) GetChildren() []NodeInterface {
	return s.children
}

// 拿到所有孩子节点的属性并整合
func (s *SyntaxNode) Attribute() string {
	if len(s.GetChildren()) == 0 {
		return s.T
	}

	attribute := ""
	for _, child := range s.GetChildren() {
		attribute = attribute + child.Attribute()
	}

	attribute += s.T

	return attribute
}
