// Hello World
package main

import (
	"encoding/json"
	"fmt"
	"regchecker/ast"
)

func main() {
	str := "a(b|c)*d"
	tokenizer := ast.NewTokenizer(str)
	tokens, _ := tokenizer.Tokenize()
	parser := ast.NewParser(tokens)
	ast, _ := parser.Parse()
	astJSON, _ := json.Marshal(ast)
	fmt.Println(string(astJSON))
}
