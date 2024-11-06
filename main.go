package main

import (
	"fmt"
	"io"
	"lexer"
	"simple_parser"
)

func main() {
	//source := "(1+(2+3))"
	//source := "if a >= 3.14"
	source := "9-5+2"
	my_lexer := lexer.NewLexer(source)
	my_parser := simple_parser.NewSimpleParser(my_lexer)
	//root, err := my_parser.Parse()
	err := my_parser.Parse()

	if err != nil && err != io.EOF {
		fmt.Println("source syntax error")
	} else if err == io.EOF {
		fmt.Println("source is legal expression")
	} else {
		//fmt.Println("unknown string")
		fmt.Println("Syntax error")
	}

	// if err == io.EOF {
	// 	fmt.Println("Syntax translation: ", root.Attribute())
	// } else {
	// 	fmt.Println("Source is syntax incorrect")
	// }
}
