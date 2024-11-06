package simple_parser

import (
	"errors"
	"fmt"
	"lexer"
)

type SimpleParser struct {
	lexer lexer.Lexer
	top   *Env
	saved *Env
}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
		top:   nil,
		saved: nil,
	}
}

func (s *SimpleParser) Parse() error {
	return s.program()
}

func (s *SimpleParser) program() error {
	s.top = nil
	return s.block()
}

func (s *SimpleParser) match(str string) error {
	if s.lexer.Lexeme != str {
		err_s := fmt.Sprintf("match error, expect %s got %s ", str, s.lexer.Lexeme)
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

	//执行语法定义的操作
	s.saved = s.top
	s.top = NewEnv(s.top)
	fmt.Print("{")

	err = s.decls()
	if err != nil {
		return err
	}

	err = s.stmts()
	if err != nil {
		return err
	}

	err = s.match("}")
	if err != nil {
		return err
	}

	//执行语法定义中的操作
	s.top = s.saved
	fmt.Print("}")
	return nil
}

func (s *SimpleParser) decls() error {
	return s.decls_r()
}

func (s *SimpleParser) decls_r() error {
	var err error
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}
	if tag.Tag == lexer.TYPE {
		//遇到int, float等变量定义字符串，因此解析变量定义
		s.lexer.ReverseScan()
		err = s.decl()
		if err != nil {
			return err
		}

		return s.decls_r()
	} else {
		//什么都不做 , 对应空
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
		str := fmt.Sprintf("in decl, expect type definition but got: %s", s.lexer.Lexeme)
		return errors.New(str)
	}
	type_str := s.lexer.Lexeme

	tag, err = s.lexer.Scan()
	if err != nil {
		return err
	}
	if tag.Tag != lexer.ID {
		str := fmt.Sprintf("in decl, expect identifier, but got: %s", s.lexer.Lexeme)
		return errors.New(str)
	}
	id_str := s.lexer.Lexeme
	//执行语法定义中的操作
	symbol := NewSymbol(id_str, type_str)
	s.top.Put(id_str, symbol)

	_, err = s.lexer.Scan()
	if err != nil {
		return err
	}

	err = s.match(";")

	return err
}

func (s *SimpleParser) stmts() error {
	/*
		消除左递归 stmts -> stmts stmt | epsilon
		stmts -> epsilon r_stmts
		r_stmts -> stmt r_stmts | epsilon
	*/

	return s.r_stmts()
}

func (s *SimpleParser) r_stmts() error {
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

		err = s.r_stmts()
	} else if tag.Tag == lexer.SEMICOLON {
		return nil
	}

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
		s.lexer.Scan()
		err = s.match(";")
		//执行语法定义的操作
		if err == nil {
			fmt.Print("; ")
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
		str := fmt.Sprintf("expect identifier , got %s ", s.lexer.Lexeme)
		return errors.New(str)
	}

	//执行语法定义的操作
	symbol := s.top.Get(s.lexer.Lexeme)

	fmt.Print(s.lexer.Lexeme)
	fmt.Print(":")
	fmt.Print(symbol.Type)

	return nil
}
