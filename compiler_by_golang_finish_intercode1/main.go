package main

import (

	//	"io"
	//	"lexer"
	//	"simple_parser"
	"fmt"
	"inter"
	"lexer"
)

func main() {
	/*
		source := "{int x; char y; {bool y; x; y;} x; y;}"
		my_lexer := lexer.NewLexer(source)
		parser := simple_parser.NewSimpleParser(my_lexer)
		err := parser.Parse()
		if err == io.EOF || err == nil {
			fmt.Println("\nparsing success")
		} else {
			fmt.Println("source is ilegal : ", err)
		}*/

	//e = (a + b) - (c + d)
	// expr_type := inter.NewType("int", lexer.TYPE, 4)
	// id_a := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "a"), expr_type)
	// id_b := inter.NewID(2, lexer.NewTokenWithString(lexer.ID, "b"), expr_type)
	// //a + b -> arith
	// arith_1, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_a, id_b)

	// id_c := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "c"), expr_type)
	// id_d := inter.NewID(2, lexer.NewTokenWithString(lexer.ID, "d"), expr_type)
	// //c + d -> arith
	// arith_2, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_c, id_d)

	// //arith: (a + b) - (c + d)
	// arith_3, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.MINUS, "-"), arith_1, arith_2)

	// id_e := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "e"), expr_type)
	// //e = (a + b) - (c + d) -> arith
	// set, _ := inter.NewSet(id_e, arith_3)

	// //生成中间代码
	// set.Gen()

	expr_type := inter.NewType("int", lexer.BASIC, 4)
	id_a := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "a"), expr_type)
	id_b := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "b"), expr_type)
	arith, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_a, id_b)

	//expr := arith.Gen()
	op := arith.Reduce()
	r := op.ToString()

	fmt.Println("\nOp node return temperay register: ", r)

}
