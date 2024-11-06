## A compiler by golang

### all tests can be passed.

### The implementation order is:

```
1. original with basic lexer and parser ->
2.  syntax_tree_node ->
3.  symbol_table ->
4.  crash_left_trace ->
5.  intercode1, intercode2 ->
```

### notion:

Before finish everything,
We has Finished necessary parts such as `lexer` and `parser`.

`original with basic lexer and parser` is put in main branch, other updated versions are put in other dev branches.
