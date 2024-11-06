我们到了简单编译器开发的最后一个阶段，也就是生成中间代码。以前我们提到过编译器分为两部分，分别为前端和后端，所谓前端就是将代码转译成中间语言，后端负责进行优化和转译成目标平台的机器指令，现在我们来到了前端的最后一个阶段。由于中间代码生成是当前所有阶段中逻辑最为复杂的部分，因此我们需要将其分解成多个容易理解的小部分，逐个击破。我们的计划是这样，首先完成比较简单的代码的中间代码生成，然后不断的提升目标代码的复杂度，然后生成更加复杂的中间代码。

本节我们要完成的如下目标代码的中间代码生成：
```
{
int a;
int b;
int c;
a = 1;
b = 2;
c = a+b;
}
```
我们要解析的代码跟前面描述的差不多，它只包含变量声明，变量赋值，和简单的加减运算，在后面我们会添加更加复杂的代码，例如if, while, for, do..while等，同时为了简单起见，我们规定变量的声明必须在代码块的起始部分，也就是不支持如下代码:
```
a = 1;
b = 2;
int c; //错误，必须在开头进行变量声明
c = a + b;
```
我们转译的中间代码就是前面说过的三地址码，它的特点是每条语句最多只能有三个变量，同时只能有一个操作符，因此在解析语句 c = a + b 时，我们首先要解析"a+b"，将它的结果放入一个临时变量，然后再将临时变量赋值给c，于是它就解析为：
```
t = a + b
c = t
```
这种计算表达式结果然后将其赋值给一个临时变量的过程我们称之为reduce。接下来我们先看要解析语言对应的语法：
```
program -> block 
block -> "{" decls stmts "}" 
decls -> decls decl | ε
decl -> type id 
type -> basic 
basic -> "int" | "float" | "bool" 
stmts -> stmts stmt | ε
stmt -> block | expr | ε
expr -> expr "+" term | expr "-" term
term -> factor 
factor -> num | id
num -> "0" | "1" | "2".... "9"
```
上面定义的语法跟我们以前接触过的没有太大区别，我们将在这个语法的基础上，在后面章节中扩展更加复杂的结构。中间代码的生成将非常依赖于语法解析树，因此我们需要在语法解析过程中构造出相应的树结构，然后再通过遍历语法树每个节点，然后根据节点的数据依次生成中间代码，下面我们先介绍节点的定义以及不同节点之间的继承关系。

首先我们定义语法树节点的接口，在项目根目录建立一个名为inter的文件夹，然后添加inter.go文件，里面的代码如下：
```
type NodeInterface interface {
	 Errors(s string)  error 
	 NewLabel() uint32 
	 EmitLabel(i uint)  
	 Emit(code string)
}
```
然后添加node.go，里面实现的是所有语法树节点的基类或者说是父节点：
```
package inter

import (
	"errors"
	 "strconv"
	 "fmt"
)

var labels uint32 //用于实现跳转的标号

type Node struct {
    lex_line  uint32 
}

func NewNode(line uint32) *Node {
	labels = 0
	return &Node {
		lex_line: line,
	}
}

func (n *Node) Errors(s string) error{
	err_s := "\nnear line " + strconv.FormatUint(uint64(n.lex_line), 10) + s
	return errors.New(err_s)
}

func (n *Node)NewLabel() uint32{
    labels = labels + 1
	return labels
}

func (n *Node) EmitLabel(i uint32) {
    fmt.Print("L" + strconv.FormatUint(uint64(i), 10) + ":")
}

func (n *Node)Emit(s string) {
	 fmt.Print("\t" + s)
}
```
所有的语法树节点都继承自Node,由于我们当前主要解析表达式，什么叫表达式呢，当一条代码执行后能产生结果，那么它就是表达式，例如类似于"a;", "a + b", "a-b"这类语句，他们都属于表达式的范畴。表达式本身涉及到一个概念叫类型，例如两个变量相加"a + b" ,其中a可能是float类型，b可能是int类型，那么两个类型进行运算时，编译器需要统一两个变量类型，通常是要把int转换为float，然后进行加法运算后结果还是float类型，所以我们需要定义和设计一个类型转换的逻辑，因此先增加一个type.go，实现代码如下：
```
package inter

import (
	"lexer"
)

type Type struct {
	width uint32  //用多少字节存储该类型
	token  *lexer.Token  //
	Lexeme  string 
}

func NewType(lexeme string, token *lexer.Token, w uint32) *Type {
	return &Type {
		width : w,
		token: token, 
		Lexeme: lexeme,
	}
}

func  Numberic(p *Type) bool {
	//查看给定类型是否属于数值类
	numberic := false
	switch p.Lexeme {
	case "int":
		numberic = true 
	case "float":
		numberic = true 
	case "char":
		numberic = true 
	}

	return numberic
}

func MaxType(p1 *Type, p2 *Type) *Type {
	/*比较类型提升，例如p1是int，p2是float, 那么就提升为float
	类型提升必须对数值类型才有效
	*/

	if Numberic(p1) == false && Numberic(p2) == false {
		return nil 
	}
	//如果两者有其一是float类型，那么就提升为float，要不然就是int
	if p1.Lexeme == "float" || p2.Lexeme == "float" {
		return NewType("float", lexer.BASIC, 8)
	} else if p1.Lexeme == "int" || p2.Lexeme == "int" {
		return NewType("int", lexer.BASIC, 4)
	} 

	return NewType("char", lexer.BASIC, 1)
}
```
这里我们将所有定义类型例如int, float, char, bool 都指定其token类型为BASIC，类型有“大小”之分，例如float 大于 int，int 大于char，因此当两个不同类型的变量进行运算时，编译器要将他们分别提升到同一类型然后才进行操作。

下面我们看看表达式节点的定义，我们先定义表达式节点的接口，因为后面有不少节点需要继承自表达式节点，在inter.go里面添加代码：
```
type ExprInterface interface {
	NodeInterface
	Gen() ExprInterface
	Reduce() ExprInterface
	Type() *Type
	ToString() string 
}
```
在接口中Gen用于生成相应语法树节点，Reduce()的作用是让当前节点生成中间代码后返回一个临时寄存器节点，具体信息在后面实现会展示。下面我们看对应表达式的节点实现，生成一个Expr.go文件，然后实现代码如下：
```
package inter
import(
	"lexer"
)

import (
	"lexer"
)

type Expr struct {
	Node  *Node 
	token   lexer.Tag 
	expr_type  *Type
}

func NewExpr(line uint32, token *lexer.Token, expr_type *Type) *Expr {
	expr := &Expr {
		Node: NewNode(line),
		token: token, 
		Type: expr_type,
	}

	return expr 
}

func (e *Expr) Errors(s string) {
	return e.Node.Errors(s)
}

func (e *Expr)NewLabel() uint32 {
	return e.Node.NewLabel()
}

func (e *Expr) EmitLabel(i uint32) {
	e.Node.EmitLabel(i)
}

func (e *Expr) Emit(code string) {
	e.Node.Emit(code)
}


func (e *Expr) Gen() ExprInterface {
	return e
}

func (e *Expr) Reduce() ExprInterface {
	return e
}

func (e *Expr) ToString() string {
    return e.token.ToString()
}

func(e *Expr)Type() *Type {
    return e.expr_type
}
```

当前我们代码会有三者表达式，一种是解析单个变量，例如语句"a;", "b;"，当我们遇到这样单变量语句时就生成一个继承于Expr的节点，其命名为ID, 如果遇到赋值语句，例如"a = 1;"，那么我们就会生成一个继承于Expr的节点，命名为Set, 如果遇到计算语句例如“a+b"，那么我们就生成一个继承于Expr的节点Arith，他们各种负责生成对应的中间代码，因此当前我们语法树节点的继承关系如下：
![请添加图片描述](https://img-blog.csdnimg.cn/958fca4cc6b14a459fb112804239b246.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAdHlsZXJfZG93bmxvYWQ=,size_14,color_FFFFFF,t_70,g_se,x_16)

下面我们看一下这几个节点的具体实现，首先是ID节点，生成id.go，添加代码如下：
```
package inter

import(
	"lexer"
)

type ID struct {
	/*
	该节点没有实现Gen,Reduce()，这意味着编译器遇到语句:"a;","b;"等时会直接越过
	不生成任何中间代码
	*/
	expr *Expr 
	Offset uint32 //相对偏移地址，用于生成中间代码，
}

func NewID(line uint32, tag lexer.Tag, expr_type *Type) *ID {
	id := &ID{
		expr: NewExpr(line, tag, expr_type),
	}

	return id 
}

func (i *ID) Errors(s string) error{
	return i.expr.Errors(s)
}

func (i *ID)NewLabel() uint32 {
	return i.expr.NewLabel()
}

func (i *ID) EmitLabel(l uint32) {
	i.expr.EmitLabel(l)
}

func (i *ID) Emit(code string) {
	i.expr.Emit(code)
}


func (i *ID) Gen() ExprInterface {
	return i 
}

func (i *ID) Reduce() ExprInterface {
	return i
}

func (i *ID) Type() *Type {
    return i.expr.Type()
}

func (i *ID) ToString() string {
	return i.expr.ToString()
}
```

我们看看Temp节点的实现，它的作用是生产一个寄存器节点，例如t, t1, t2这些，创建文件temp.go,其代码如下：
```
package inter

import(
	"lexer"
	"strconv"
)
/*
Temp节点表示中间代码中的临时寄存器变量
*/
type Temp struct {
	expr *Expr
	number uint32 
}

var count uint32 

func NewTemp(line uint32, expr_type *Type) *Temp {
	token := lexer.NewToken(lexer.TEMP)
	count = count + 1
	
    temp := &Temp {
		expr: NewExpr(line,  &token, expr_type),
		number : count,
	}

   return temp
}

func (t *Temp) Errors(s string) error {
	return t.expr.Errors(s)
}

func (t *Temp) NewLabel() uint32 {
	return t.expr.NewLabel()
}

func (t *Temp)EmitLabel(i uint32) {
	 t.expr.EmitLabel(i)
}

func (t *Temp)Emit(code string)  {
    t.expr.Emit(code)
}

func (t *Temp) Gen() ExprInterface {
	return t
}

func (t *Temp) Reduce() ExprInterface {
	return t
}

func (t *Temp) ToString() string{
    return "t" + strconv.FormatUint(uint64(number), 10)
}

func (t *Temp) Type() *Type {
	return t.expr.Type()
}
```
Temp节点对应中间代码中的临时寄存器变量，它的作用在后面代码实现中间代码生成时就会明晰，接下来我们看看Op节点的实现，创建Op.go，实现代码如下：
```
package inter

import (
	"lexer"
)

type Op struct {
	expr *Expr 
	child ExprInterface
	line uint32
	expr_type *Type
}

func NewOp(line uint32, token *lexer.Token, expr_type *Type) *Op{
	op := &Op {
		expr: NewExpr(line, token, expr_type),
		child: nil,
		line: line,
		expr_type: expr_type,
	}

	return op 
}

func (o *Op) Errors(s string) error {
	return o.expr.Errors(s)
}

func (o *Op) NewLabel() uint32 {
	return o.expr.NewLabel()
}

func (o *Op) EmitLabel(i uint32) {
    o.expr.EmitLabel(i)
}

func (o *Op) Emit(code string) {
	o.expr.Emit(code)
}

func (o *Op) Gen() ExprInterface{
    return o 
}

func (o *Op) Reduce() ExprInterface {
	if o.child != nil {
		/*调用子节点的Gen函数，让子节点先生成中间代码,
		子节点生成中间代码后会返回一个Expr节点，然后这里将返回的节点赋值给
		一个临时寄存器变量

		具体逻辑为当编译器遇到语句 a + b 就会生成Op节点,
	    那么a + b对应一个Arith节点，它对应child对象，
		执行child.Gen()会生成中间代码对应的字符串:
		a + b
	    接下来我们创建一个临时寄存器变量例如t2,然后生成语句
		t2 = a + b
		*/
		x := o.child.Gen()
		t := NewTemp(o.line, o.expr_type)
		o.expr.Emit(t.ToString() + "=" + x.ToString())
	}

	return 0
}

func (o *Op) ToString() string {
	return o.expr.ToString()
}


```
代码实现中有一点需要注意就是它的Reduce()函数，当编译器遇到语句"a+b"或者"a-b"时就会生成一个Op节点，同时创建对应的子Arith节点，在生成中间代码时，先调用Arith的Gen函数生成代码字符串" a + b" 或是 "a - b"，然后在创建一个临时寄存器变量t,最后生成中间代码:
```
t = a + b
```
或者是:
```
t = a - b
```
下面我们看节点Arith的实现，当编译器读到"a+b"这类语句时除了创建Op节点外也会创建Arith节点，它负责将操作 a op b 转换成指令 a op b，这类的op可以对应+,-.\*, \，创建arith.go，实现代码如下：
```
package inter

import (
	"lexer"
)



type Arith struct {
	op *Op
	line uint32
	token *lexer.Token
	expr1 ExprInterface
	expr2 ExprInterface
	expr_type *Type
}

func NewArith(line uint32, token *lexer.Token, expr1 *Expr, 
	expr2 *Expr) *Arith {
	arith := &Arith{
		/*
		暂时还不能决定类型，因为两个表达式expr1, expr2各自的类型不一样，需要解析他们的类型，然后
		进行类型提升后才能决定当前节点的类型
		*/
		op: NewOp(line, token, nil),
		line : line ,
		token: token ,
		expr1 : expr1, 
		expr2 : expr2,
		expr_type: expr_type,
	}

	arith.op.child = arith 
    if expr_type == nil {
		//表达式的类型不能为空
		return nil, arith.op.Errors("type error")
	}

	return arith, nil
}

func (a *Arith) Errors(s string) error {
	return a.op.Errors(s)
}

func (a *Arith) NewLabel() uint32 {
	return a.op.NewLabel()
}

func (a *Arith) EmitLabel(i uint32) {
	a.op.EmitLabel(i)
}

func (a *Arith) Emit(code string) {
	a.op.Emit(code)
}

func (a *Arith) Gen() ExprInterface{
    /*
	我们可能会遇到复杂的组合表达式例如 (a+b) + (c+d)，
	于是expr1 对应a+b, expr2 对应c+d，
	此时节点生成中间代码时，需要先让expr1和expr2生成代码
	*/
    return NewArith(a.line, a.token, a.expr1.Reduce(), a.expr2.Reduce())
}

func (a *Arith) Reduce() ExprInterface {
	return a.op.Reduce()
}

func (a *Arith) ToString() string {
	return a.expr1.ToString() + " " + a.token.ToString() + " " + a.expr2.ToString()
}

func (a *Arith) Type() *Type {
    return a.expr_type
}
```

Arith节点的逻辑不容易理解，假设给定语句"a + (b+c)"，当编译器遇到它时就会生成Arith节点，它会把其理解为两部分，首先它把a当做ID节点，这里对应expr1，然后把b+c当做Arith节点，这里对应expr2,在Arith的Gen函数中，它分别调用expr1.Reduce()和expr2.Reduce()，由于expr1是ID节点，它的Reduce函数会调用包含在ID节点中的expr对象的Reduce函数，后者的Reduce函数根据expr.go里面的实现就会返回它自己作为ExprInterface接口的实例。注意到节点Expr的ToString()接口会返回其对应token对象的字符串，由于当前token对应的字符串恰好就是变量a对应的字符串也就是"a"。

expr2对应的就是Arith节点，它里面两个expr1和expr2分别对应两个ID节点分别为b 和 c，两个ID节点的Reduce()接口其实直接调用了他们内部expr对象的Reduce()接口，于是返回的也是ID节点内部的expr对象，于是expr1.Reduc()执行后返回它对应内部的expr对象，这个对象的ToString()就会返回字符串"b"，expr2.Reduce()执行后会返回它对应内部的expr对象,这个对象的ToString就会返回字符串"c"，于是这个Arith节点的ToString接口调用后就会返回字符串"b+c"。

这部分的逻辑比较复杂，最好的理解方式还是看对应视频的调试演示，仅仅通过文字叙述很难说清楚，当然也需要你慢慢磨代码。为了更好的理解当前完成的代码，我们运行当前已经完成的代码看看，在main.go中输入代码如下：
```
package main

import (
	"fmt"
	"inter"
	"lexer"
)

func main() {
	
	expr_type := inter.NewType("int", lexer.BASIC, 4)
	id_a := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "a"), expr_type)
	id_b := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "b"), expr_type)
	arith, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_a, id_b)

	//expr := arith.Gen()
	op := arith.Reduce()
	r := op.ToString()

	fmt.Printf("\nOp node return temperay register: ", r)
}
```
上面代码运行后所得结果如下：
```
t1 = a + b
Op node return temperay register:  t1
```
从输出看，我们手动构造的节点能输出正确的中间代码，对于语句 a + b，它能够先分配一个临时寄存器变量t1,然后将a+b的结果赋值给寄存器变量t1，本节的逻辑比较复杂，更好的理解本节代码逻辑，请在B站搜索Coding迪斯尼，查看调试演示视频。
