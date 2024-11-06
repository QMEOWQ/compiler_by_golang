//backus-naur form
//list -> "(" list ")"
//list -> list "+" list
//list -> list "-" list
//list -> number
//number -> ...

// program ->  block  {top = nil}
// block -> '{'  decls stmts '}'  {saved = top; top = NewEnv(top); print("{"}
// decls -> decls decl | ε
// decl -> type id ";"  {s = NewSymbol(type.lexeme, id.lexeme); top[id.lexeme] = s}
// stmts -> stmts stmt | ε
// stmt -> block | factor ";"  {print(";")}
// factor -> id  {s = top[id.lexeme]; print(id.lexeme); print(";"); print(s.type.lexeme);}

//crash left trace
// decls -> decls_r
// decls_r -> decl decls_r | ε

// stmts -> stmts stmt
// stmts-> stmts_r
// stmts_r -> stmt stmts_r | ε

package simple_parser

import (
	"errors"
	"fmt"
	"lexer"
)

type SimpleParser struct {
	lexer lexer.Lexer
	top   *Env //当前作用域的符号表
	saved *Env //进入下一个作用域时，记录当前作用域符号表
}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
		top:   nil,
		saved: nil,
	}
}

func (s *SimpleParser) Parse() error {
	//进入program
	return s.program()
}

func (s *SimpleParser) program() error {
	//program -> block {top = nil}
	s.top = nil
	return s.block()
}

func (s *SimpleParser) match(str string) error {
	if s.lexer.Lexeme != str {
		err_s := fmt.Sprintf("match error, expect %s, but get %s", str, &s.lexer.Lexeme)
		return errors.New(err_s)
	}

	return nil
}

func (s *SimpleParser) block() error {
	s.lexer.Scan()
	err := s.match("{")
	if err != nil {
		return err
	}

	//执行语法说明操作
	s.saved = s.top
	s.top = NewEnv(s.top)
	fmt.Printf("{")

	//decls parse
	err = s.decls()
	if err != nil {
		return err
	}

	//stmas parse
	err = s.stmts()
	if err != nil {
		return err
	}

	//match "}"
	err = s.match("}")
	if err != nil {
		return err
	}

	//执行语法定义操作, 一个作用域结束
	s.top = s.saved
	fmt.Printf("}")

	return nil
}

func (s *SimpleParser) decls() error {
	//
	//crash left trace
	return s.decls_r()
}

func (s *SimpleParser) decls_r() error {
	var err error
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag == lexer.TYPE {
		s.lexer.ReverseScan()
		err = s.decl()
		if err != nil {
			return err
		}
		return s.decls_r()
	} else {
		//
		s.lexer.ReverseScan()
	}

	return nil
}

func (s *SimpleParser) decl() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag != lexer.TYPE {
		str := fmt.Sprintf("in decl, expect type but got: %s", s.lexer.Lexeme)
		return errors.New(str)
	}

	type_str := s.lexer.Lexeme

	tag, err = s.lexer.Scan()
	if err != nil {
		return err
	}
	if tag.Tag != lexer.ID {
		str := fmt.Sprintf("in decl, expect identifier but got: %s", s.lexer.Lexeme)
		return errors.New(str)
	}

	id_str := s.lexer.Lexeme
	//执行语法定义操作
	symbol := NewSymbol(id_str, type_str)
	s.top.Put(id_str, symbol)

	_, err = s.lexer.Scan()
	if err != nil {
		return err
	}

	err = s.match(";")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf(";")

	return err
}

func (s *SimpleParser) stmts() error {
	//
	//crash left trace
	return s.stmts_r()
}

func (s *SimpleParser) stmts_r() error {
	// stmts_r -> stmt stmts_r | ε
	//crash left trace
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag == lexer.ID || tag.Tag == lexer.LEFT_BRACE {
		s.lexer.ReverseScan()
		err = s.stmt()
		if err != nil {
			return err
		}
		err = s.stmts_r()
	} else if tag.Tag == lexer.SEMICOLON {
		// ; 也是一个表达式
		return nil
	}

	//return err
	return nil
}

func (s *SimpleParser) stmt() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag == lexer.LEFT_BRACE {
		s.lexer.ReverseScan()
		err = s.block()
	} else if tag.Tag == lexer.ID {
		s.lexer.ReverseScan()
		err = s.factor()
		if err != nil {
			return err
		}
		s.lexer.Scan()
		err = s.match(";")
		if err == nil {
			fmt.Print(";")
		}
	} else {
		err = errors.New("stmt parsing error")
	}

	return err
}

func (s *SimpleParser) factor() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag != lexer.ID {
		str := fmt.Sprintf("expect identifier, but got: %s", s.lexer.Lexeme)
		return errors.New(str)
	}

	symbol := s.top.Get(s.lexer.Lexeme)

	//如果symbol为空，说明代码出现引用未定义变量的错误
	if symbol == nil {
		return errors.New("undefined variable")
	}

	//eg. x : int
	fmt.Print(s.lexer.Lexeme)
	fmt.Print(" : ")
	fmt.Print(symbol.Type)

	return nil
}
