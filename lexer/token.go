package lexer

type Tag uint32

const (
	// AND Tag = iota + 256 // &&
	// BASIC
	// BREAK // break
	// DO
	// EQ               // ==
	// FALSE            // false
	// GE               // >=
	// LE               // <=
	// ID               // identifier
	// IF               // if
	// ELSE             // else
	// MINUS            // -
	// PLUS             // +
	// NE               // !=
	// NUM              // eg. 134, 56...
	// OR               // ||
	// REAL             // eg. 3.14, 2.71...
	// STRING           // string
	// TRUE             // true
	// WHILE            // while
	// LEFT_BRACE       // {
	// RIGHT_BRACE      // }
	// LEFT_BRACKET     // (
	// RIGHT_BRACKET    // )
	// AND_OPERATOR     // &
	// OR_OPERATOR      // |
	// ASSIGN_OPERATOR  // =
	// NEGATE_OPERATOR  // !
	// LESS_OPERATOR    // <
	// GREATER_OPERATOR // >

	// EOF   // end of file
	// ERROR // error

	AND Tag = iota + 256
	BASIC
	BREAK
	DO
	EQ
	FALSE
	GE
	ID
	IF
	ELSE
	INDEX
	LE
	INT
	FLOAT
	MINUS
	PLUS
	NE
	NUM
	OR
	REAL
	TEMP
	TRUE
	WHILE
	LEFT_BRACE    // "{"
	RIGHT_BRACE   // "}"
	LEFT_BRACKET  //"("
	RIGHT_BRACKET //")"
	AND_OPERATOR
	OR_OPERATOR
	ASSIGN_OPERATOR
	NEGATE_OPERATOR
	LESS_OPERATOR
	GREATER_OPERATOR
	TYPE      //对应int , float, bool, char 等类型定义
	SEMICOLON // ;

	EOF

	ERROR
)

var token_map = make(map[Tag]string)

func init() {
	// token_map[AND] = "&&"
	// token_map[BASIC] = "BASIC"
	// token_map[BREAK] = "break"
	// token_map[DO] = "do"
	// token_map[EQ] = "=="
	// token_map[FALSE] = "false"
	// token_map[GE] = ">="
	// token_map[LE] = "<="
	// token_map[ID] = "identifier"
	// token_map[IF] = "if"
	// token_map[ELSE] = "else"
	// token_map[MINUS] = "-"
	// token_map[PLUS] = "+"
	// token_map[NE] = "!="
	// token_map[NUM] = "NUM"
	// token_map[OR] = "OR"
	// token_map[REAL] = "REAL"
	// token_map[STRING] = "STRING"
	// token_map[TRUE] = "TRUE"
	// token_map[WHILE] = "while"
	// token_map[LEFT_BRACE] = "{"
	// token_map[RIGHT_BRACE] = "}"
	// token_map[AND_OPERATOR] = "&"
	// token_map[OR_OPERATOR] = "|"
	// token_map[ASSIGN_OPERATOR] = "="
	// token_map[NEGATE_OPERATOR] = "!"
	// token_map[LESS_OPERATOR] = "<"
	// token_map[GREATER_OPERATOR] = ">"
	// token_map[EOF] = "end of file"
	// token_map[ERROR] = "error"
	// token_map[LEFT_BRACKET] = "("
	// token_map[RIGHT_BRACKET] = ")"

	token_map[AND] = "&&"
	token_map[BASIC] = "BASIC"
	token_map[DO] = "do"
	token_map[ELSE] = "else"
	token_map[EQ] = "EQ"
	token_map[FALSE] = "FALSE"
	token_map[GE] = "GE"
	token_map[ID] = "ID"
	token_map[IF] = "if"
	token_map[INT] = "int"
	token_map[FLOAT] = "float"
	token_map[INDEX] = "INDEX"
	token_map[LE] = "<="
	token_map[MINUS] = "-"
	token_map[PLUS] = "+"
	token_map[NE] = "!="
	token_map[NUM] = "NUM"
	token_map[OR] = "OR"
	token_map[REAL] = "REAL"
	token_map[TEMP] = "TEMP"
	token_map[TRUE] = "TRUE"
	token_map[WHILE] = "while"
	token_map[AND_OPERATOR] = "&"
	token_map[OR_OPERATOR] = "|"
	token_map[ASSIGN_OPERATOR] = "="
	token_map[NEGATE_OPERATOR] = "!"
	token_map[LESS_OPERATOR] = "<"
	token_map[GREATER_OPERATOR] = ">"
	token_map[LEFT_BRACE] = "{"
	token_map[RIGHT_BRACE] = "}"
	token_map[LEFT_BRACKET] = "("
	token_map[RIGHT_BRACKET] = ")"
	token_map[EOF] = "EOF"
	token_map[ERROR] = "ERROR"
	token_map[TYPE] = "TYPE"
	token_map[SEMICOLON] = ";"

}

type Token struct {
	Tag Tag
} // Token represents a single token in the input program.

func (t *Token) ToString() string {
	return token_map[t.Tag]
}

func NewToken(tag Tag) Token {
	return Token{
		Tag: tag,
	}
}
