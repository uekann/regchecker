package ast_test

import (
	"regchecker/ast"
	"testing"
)

var astStringTests = []struct {
	ast ast.AST
	str string
}{
	{ast.AST{Kind: ast.ASTKindChar, Value: "a"}, "a"},
	{ast.AST{Kind: ast.ASTKindUnion, Children: []ast.AST{
		{Kind: ast.ASTKindChar, Value: "a"},
		{Kind: ast.ASTKindChar, Value: "b"},
	}}, "a|b"},
	{ast.AST{Kind: ast.ASTKindConcat, Children: []ast.AST{
		{Kind: ast.ASTKindChar, Value: "a"},
		{Kind: ast.ASTKindChar, Value: "b"},
	}}, "ab"},
	{ast.AST{Kind: ast.ASTKindStar, Children: []ast.AST{
		{Kind: ast.ASTKindChar, Value: "a"},
	}}, "a*"},
	{ast.AST{Kind: ast.ASTKindGroup, Children: []ast.AST{
		{Kind: ast.ASTKindChar, Value: "a"},
	}}, "(a)"},
}

func TestASTString(t *testing.T) {
	for _, tt := range astStringTests {
		if got, want := tt.ast.String(), tt.str; got != want {
			t.Errorf("ast.String() = %q, want %q", got, want)
		}
	}
}
