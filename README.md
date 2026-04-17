# 编译器项目（Go）

这是一个使用 Go 实现的简化编译器项目，当前以教学和工程演进为目标，核心流程为：

`源代码 -> 词法分析(Lexer) -> 语法分析(Parser) -> 中间代码生成(Inter)`

项目已开始扩展 `golex` 生成器子系统，用于后续按视频标题推进词法生成器相关能力。

## 模块结构

```text
compiler_by_golang_brandnew/
├── main.go                 # 入口示例
├── go.mod                  # 单模块依赖管理
├── lexer/                  # 词法分析
├── parser/                 # 语法分析与符号环境
├── inter/                  # 中间代码生成
├── golex/                  # 词法生成器（持续实现中）
│   ├── spec/
│   ├── reader/
│   ├── regex/
│   ├── nfa/
│   ├── dfa/
│   ├── minimize/
│   ├── codegen/
│   └── runtime/
└── testdata/               # 样例输入
```

## 核心能力

- Lexer：将字符流扫描为 Token 流，支持关键字识别与回退。
- Parser：递归下降解析，维护作用域链符号表，构建表达式语义结构。
- Inter：基于节点接口生成三地址码，支持临时变量与标签。

## 示例语法（当前阶段）

```text
program  -> block
block    -> "{" decls stmts "}"
decls    -> decls decl | ε
decl     -> type id ";"
stmts    -> stmts stmt | ε
stmt     -> assign | expr ";"
assign   -> id "=" expr ";"
expr     -> expr "+" term | expr "-" term | term
term     -> factor
factor   -> num | real | id | "(" expr ")"
```

## 快速开始

```bash
go test ./...
go run .
```

## 运行输出特征

- 中间代码以标签和三地址码形式输出。
- 复杂表达式会被拆分并引入临时变量（如 `t1`, `t2`）。

## 规划路线

后续将按视频标题分阶段实现：

- `lexReader`：规则文件读取和分段解析
- `RegParser`：正则表达式解析
- `NFA/DFA`：状态机构建与转换
- `DFA 最小化`：分区优化
- `代码生成`：输出可执行扫描器

详细规划见：`docs/video-roadmap.md`（已在 `.gitignore` 中忽略，不参与上传）。
