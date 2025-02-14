package ast_test

import (
	"errors"
	"regchecker/ast"
	"testing"
)

var astErrorTests = []struct {
	str string
	err ast.ParseErrorKind
}{
	{"", ast.ParseErrorKindUnexpectedEOF},
	{"a)", ast.ParseErrorKindParenNotOpened},
	{"*", ast.ParseErrorKindUnexpectedStar},
	{"(a", ast.ParseErrorKindParenNotClosed},
	{"a**", ast.ParseErrorKindUnexpectedStar},
	{"a|*", ast.ParseErrorKindUnexpectedStar},
	{"(a**)", ast.ParseErrorKindUnexpectedStar},
}

func equalASTs(a, b ast.AST) bool {
	if a.Kind != b.Kind || a.Value != b.Value {
		return false
	}
	if len(a.Children) != len(b.Children) {
		return false
	}
	for i := range a.Children {
		if !equalASTs(*a.Children[i], *b.Children[i]) {
			return false
		}
	}
	return true
}

func TestParse(t *testing.T) {
	for _, tt := range astStringTests {
		tokenizer := ast.NewTokenizer(tt.str)
		tokens, err := tokenizer.Tokenize()
		if err != nil {
			t.Errorf("Tokenize(%q) returned error: %v", tt.str, err)
			continue
		}
		parser := ast.NewParser(tokens)
		ast, err := parser.Parse()
		if err != nil {
			t.Errorf("Parse(%q) returned error: %v", tt.str, err)
			continue
		}
		if !equalASTs(*ast, tt.ast) {
			t.Errorf("Parse(%q) = %v, want %v", tt.str, ast, tt.ast)
		}
	}
}

func TestParseError(t *testing.T) {
	for _, tt := range astErrorTests {
		// tokens := ast.Tokenize(tt.str)
		tokenizer := ast.NewTokenizer(tt.str)
		tokens, err := tokenizer.Tokenize()
		if err != nil {
			t.Errorf("Tokenize(%q) returned error: %v", tt.str, err)
			continue
		}
		parser := ast.NewParser(tokens)
		_, err = parser.Parse()
		if err == nil {
			t.Errorf("Parse(%q) did not return error", tt.str)
			continue
		}
		var parseError *ast.ParseError
		if errors.As(err, &parseError) {
			if parseError.Kind != tt.err {
				t.Errorf("Parse(%q) returned error %v, want %v", tt.str, parseError.Kind, tt.err)
			}
		}
	}
}
