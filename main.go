package main

import (
	"dragon-compiler/lexer"
	simpleparser "dragon-compiler/parser"
	"fmt"
	//"inter"
)

func main() {

	// expr_type := inter.NewType("int", lexer.BASIC, 4)
	// id_a := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "a"), expr_type)
	// id_b := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "b"), expr_type)
	// //a+b
	// arith1, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_a, id_b)

	// id_c := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "c"), expr_type)
	// id_d := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "d"), expr_type)
	// arith2, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_c, id_d)

	// arith3, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.MINUS, "-"), arith1, arith2)

	// //arith3.Reduce()

	// id_e := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "e"), expr_type)
	// //e = (a+b) - (b+c) -> c = t1 - t2
	// set, _ := inter.NewSet(id_e, arith3)
	// set.Gen()

	// 示例1：基本算术运算
	source1 := `{int x; float y ; float c; float d;
	              x = 1; y = 3.14;
	              c = x + y;
	              d = x + y + c;
	              }`

	fmt.Println("=== 示例1：基本算术运算 ===")
	my_lexer1 := lexer.NewLexer(source1)
	parser1 := simpleparser.NewSimpleParser(my_lexer1)
	parser1.Parse()

	// 示例2：更复杂的表达式
	source2 := `{int a; int b; int c; int result;
              a = 10; b = 20; c = 30;
              result = a + b - c;
              }`

	fmt.Println("\n=== 示例2：复杂表达式 ===")
	my_lexer2 := lexer.NewLexer(source2)
	parser2 := simpleparser.NewSimpleParser(my_lexer2)
	parser2.Parse()

}
