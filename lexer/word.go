package lexer

type Word struct {
	//eg. int abc = 123
	lexeme string // "abc"
	Tag    Token  // 123 -> ID
}

func NewWordToken(s string, tag Tag) Word {
	return Word{
		lexeme: s,
		Tag:    NewToken(tag),
	}
}

func GetKeyWords() []Word {
	// key_words := []Word{}
	// key_words = append(key_words, NewWordToken("&&", AND))
	// key_words = append(key_words, NewWordToken("||", OR))
	// key_words = append(key_words, NewWordToken("!", NEGATE_OPERATOR))
	// key_words = append(key_words, NewWordToken("==", EQ))
	// key_words = append(key_words, NewWordToken("!=", NE))
	// key_words = append(key_words, NewWordToken(">=", GE))
	// key_words = append(key_words, NewWordToken("<=", LE))
	// key_words = append(key_words, NewWordToken(">", GREATER_OPERATOR))
	// key_words = append(key_words, NewWordToken("<", LESS_OPERATOR))
	// key_words = append(key_words, NewWordToken("=", ASSIGN_OPERATOR))
	// key_words = append(key_words, NewWordToken("+", PLUS))
	// key_words = append(key_words, NewWordToken("-", MINUS))
	// key_words = append(key_words, NewWordToken("NUM", NUM))
	// key_words = append(key_words, NewWordToken("REAL", REAL))
	// key_words = append(key_words, NewWordToken("identifier", ID))
	// key_words = append(key_words, NewWordToken("STRING", STRING))
	// key_words = append(key_words, NewWordToken("OR", OR))
	// key_words = append(key_words, NewWordToken("|", OR_OPERATOR))
	// // key_words = append(key_words, NewWordToken("*", MUL))
	// // key_words = append(key_words, NewWordToken("/", DIV))
	// // key_words = append(key_words, NewWordToken("%", MOD))
	// // key_words = append(key_words, NewWordToken("(", LPAREN))
	// // key_words = append(key_words, NewWordToken(")", RPAREN))
	// key_words = append(key_words, NewWordToken("{", LEFT_BRACE))
	// key_words = append(key_words, NewWordToken("}", RIGHT_BRACE))
	// // key_words = append(key_words, NewWordToken("[", LBRACKET))
	// // key_words = append(key_words, NewWordToken("]", RBRACKET))
	// key_words = append(key_words, NewWordToken("if", IF))
	// key_words = append(key_words, NewWordToken("else", ELSE))
	// key_words = append(key_words, NewWordToken("while", WHILE))
	// key_words = append(key_words, NewWordToken("do", DO))
	// key_words = append(key_words, NewWordToken("break", BREAK))
	// key_words = append(key_words, NewWordToken("end of file", EOF))
	// key_words = append(key_words, NewWordToken("error", ERROR))

	key_words := []Word{}
	key_words = append(key_words, NewWordToken("&&", AND))
	key_words = append(key_words, NewWordToken("||", OR))
	key_words = append(key_words, NewWordToken("==", EQ))
	key_words = append(key_words, NewWordToken("!=", NE))
	key_words = append(key_words, NewWordToken("<=", LE))
	key_words = append(key_words, NewWordToken(">=", GE))
	key_words = append(key_words, NewWordToken("minus", MINUS))
	key_words = append(key_words, NewWordToken("true", TRUE))
	key_words = append(key_words, NewWordToken("false", FALSE))
	key_words = append(key_words, NewWordToken("t", TEMP))
	key_words = append(key_words, NewWordToken("if", IF))
	key_words = append(key_words, NewWordToken("else", ELSE))

	//添加类型定义
	key_words = append(key_words, NewWordToken("int", TYPE))
	key_words = append(key_words, NewWordToken("float", TYPE))
	key_words = append(key_words, NewWordToken("bool", TYPE))
	key_words = append(key_words, NewWordToken("char", TYPE))

	return key_words
}

func (w *Word) ToString() string {
	return w.lexeme
}
