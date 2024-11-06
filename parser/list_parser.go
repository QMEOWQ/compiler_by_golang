//backus-naur form
//list -> "(" list ")"
//list -> list "+" list
//list -> list "-" list
//list -> number
//number -> ...

package simple_parser

import (
	"errors"
	"fmt"
	"lexer"
)

type SimpleParser struct {
	lexer lexer.Lexer
}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
	}
}

func (s *SimpleParser) Parse() error {
	return s.list()
}

func (s *SimpleParser) digit() error {
	//digit -> "0" {print("0")} | ... | "9" {print("9")}
	token, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if token.Tag != lexer.NUM || len(s.lexer.Lexeme) > 1 {
		s := fmt.Sprintf("parsing error, expect a single digit but got %s", &s.lexer.Lexeme)
		return errors.New(s)
	}

	fmt.Print(s.lexer.Lexeme)

	return nil
}

func (s *SimpleParser) rest() error {
	//rest -> "+" digit {print("+")} rest | "-" digit {print("-")} rest | ε
	token, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if token.Tag == lexer.PLUS {
		err = s.digit()
		if err != nil {
			return err
		}

		fmt.Print("+")

		err = s.rest()
		if err != nil {
			return err
		}

	} else if token.Tag == lexer.MINUS {
		err = s.digit()
		if err != nil {
			return err
		}

		fmt.Print("-")

		err = s.rest()
		if err != nil {
			return err
		}
	} else {
		// ε
	}

	return err
}

// 新的list，消除左递归
func (s *SimpleParser) list() error {
	//list -> digit rest
	err := s.digit()
	if err != nil {
		return err
	}

	err = s.rest()

	return err
}
