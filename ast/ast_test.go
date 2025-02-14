package ast_test

import (
	"regchecker/ast"
	"testing"
)

var astStringTests = []struct {
	ast ast.AST
	str string
}{
	// a
	{ast.AST{Kind: ast.ASTKindChar, Value: "a"}, "a"},

	// a|b
	{ast.AST{Kind: ast.ASTKindUnion, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
		&ast.AST{Kind: ast.ASTKindChar, Value: "b"},
	}}, "a|b"},

	// ab
	{ast.AST{Kind: ast.ASTKindConcat, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
		&ast.AST{Kind: ast.ASTKindChar, Value: "b"},
	}}, "ab"},

	// a*
	{ast.AST{Kind: ast.ASTKindStar, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
	}}, "a*"},

	// (a)
	{ast.AST{Kind: ast.ASTKindGroup, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
	}}, "(a)"},

	// ab*
	{ast.AST{Kind: ast.ASTKindConcat, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
		&ast.AST{Kind: ast.ASTKindStar, Children: []*ast.AST{
			&ast.AST{Kind: ast.ASTKindChar, Value: "b"},
		}},
	}}, "ab*"},

	// a|b*
	{ast.AST{Kind: ast.ASTKindUnion, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
		&ast.AST{Kind: ast.ASTKindStar, Children: []*ast.AST{
			&ast.AST{Kind: ast.ASTKindChar, Value: "b"},
		}},
	}}, "a|b*"},

	// ab|c
	{ast.AST{Kind: ast.ASTKindUnion, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindConcat, Children: []*ast.AST{
			&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
			&ast.AST{Kind: ast.ASTKindChar, Value: "b"},
		}},
		&ast.AST{Kind: ast.ASTKindChar, Value: "c"},
	}}, "ab|c"},

	// (a|b)*
	{ast.AST{Kind: ast.ASTKindGroup, Children: []*ast.AST{
		&ast.AST{Kind: ast.ASTKindStar, Children: []*ast.AST{
			&ast.AST{Kind: ast.ASTKindGroup, Children: []*ast.AST{
				&ast.AST{Kind: ast.ASTKindUnion, Children: []*ast.AST{
					&ast.AST{Kind: ast.ASTKindChar, Value: "a"},
					&ast.AST{Kind: ast.ASTKindChar, Value: "b"},
				}},
			}},
		}},
	}}, "((a|b)*)"},
}

func TestASTString(t *testing.T) {
	for _, tt := range astStringTests {
		if got, want := tt.ast.String(), tt.str; got != want {
			t.Errorf("ast.String() = %q, want %q", got, want)
		}
	}
}
