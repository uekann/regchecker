package ast_test

import (
	"regchecker/ast"
	"testing"
)

var tokenTests = []struct {
	s      string
	tokens []ast.Token
}{
	{"a", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a|b", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindPipe, Value: "|"},
		{Kind: ast.TokenKindChar, Value: "b"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a|b*", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindPipe, Value: "|"},
		{Kind: ast.TokenKindChar, Value: "b"},
		{Kind: ast.TokenKindStar, Value: "*"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a|b*c", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindPipe, Value: "|"},
		{Kind: ast.TokenKindChar, Value: "b"},
		{Kind: ast.TokenKindStar, Value: "*"},
		{Kind: ast.TokenKindChar, Value: "c"},
		{Kind: ast.TokenKindEOF},
	}},
	{"(a|b)*c", []ast.Token{
		{Kind: ast.TokenKindLParen, Value: "("},
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindPipe, Value: "|"},
		{Kind: ast.TokenKindChar, Value: "b"},
		{Kind: ast.TokenKindRParen, Value: ")"},
		{Kind: ast.TokenKindStar, Value: "*"},
		{Kind: ast.TokenKindChar, Value: "c"},
		{Kind: ast.TokenKindEOF},
	}},
	{"^a", []ast.Token{
		{Kind: ast.TokenKindCaret, Value: "^"},
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a$", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindDollar, Value: "$"},
		{Kind: ast.TokenKindEOF},
	}},
	{"\\a", []ast.Token{
		{Kind: ast.TokenKindEscape, Value: "\\a"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a{1,2}", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindLBrace, Value: "{"},
		{Kind: ast.TokenKindNumber, Value: "1"},
		{Kind: ast.TokenKindComma, Value: ","},
		{Kind: ast.TokenKindNumber, Value: "2"},
		{Kind: ast.TokenKindRBrace, Value: "}"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a{1234,5678}", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindLBrace, Value: "{"},
		{Kind: ast.TokenKindNumber, Value: "1234"},
		{Kind: ast.TokenKindComma, Value: ","},
		{Kind: ast.TokenKindNumber, Value: "5678"},
		{Kind: ast.TokenKindRBrace, Value: "}"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a{1}", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindLBrace, Value: "{"},
		{Kind: ast.TokenKindNumber, Value: "1"},
		{Kind: ast.TokenKindRBrace, Value: "}"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a{1,}", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindLBrace, Value: "{"},
		{Kind: ast.TokenKindNumber, Value: "1"},
		{Kind: ast.TokenKindComma, Value: ","},
		{Kind: ast.TokenKindRBrace, Value: "}"},
		{Kind: ast.TokenKindEOF},
	}},
	{"a{,2}", []ast.Token{
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindLBrace, Value: "{"},
		{Kind: ast.TokenKindComma, Value: ","},
		{Kind: ast.TokenKindNumber, Value: "2"},
		{Kind: ast.TokenKindRBrace, Value: "}"},
		{Kind: ast.TokenKindEOF},
	}},
	{"[a-z]", []ast.Token{
		{Kind: ast.TokenKindLBracket, Value: "["},
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindHyphen, Value: "-"},
		{Kind: ast.TokenKindChar, Value: "z"},
		{Kind: ast.TokenKindRBracket, Value: "]"},
		{Kind: ast.TokenKindEOF},
	}},
	{"[a-z0-9]", []ast.Token{
		{Kind: ast.TokenKindLBracket, Value: "["},
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindHyphen, Value: "-"},
		{Kind: ast.TokenKindChar, Value: "z"},
		{Kind: ast.TokenKindChar, Value: "0"},
		{Kind: ast.TokenKindHyphen, Value: "-"},
		{Kind: ast.TokenKindChar, Value: "9"},
		{Kind: ast.TokenKindRBracket, Value: "]"},
		{Kind: ast.TokenKindEOF},
	}},
	{"[^a-z]", []ast.Token{
		{Kind: ast.TokenKindLBracket, Value: "["},
		{Kind: ast.TokenKindCaret, Value: "^"},
		{Kind: ast.TokenKindChar, Value: "a"},
		{Kind: ast.TokenKindHyphen, Value: "-"},
		{Kind: ast.TokenKindChar, Value: "z"},
		{Kind: ast.TokenKindRBracket, Value: "]"},
		{Kind: ast.TokenKindEOF},
	}},
	{"[\\n\\-]", []ast.Token{
		{Kind: ast.TokenKindLBracket, Value: "["},
		{Kind: ast.TokenKindEscape, Value: "\\n"},
		{Kind: ast.TokenKindEscape, Value: "\\-"},
		{Kind: ast.TokenKindRBracket, Value: "]"},
		{Kind: ast.TokenKindEOF},
	}},
}

func equalTokens(a, b []ast.Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Kind != b[i].Kind || a[i].Value != b[i].Value {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	for _, tt := range tokenTests {
		tokenizer := ast.NewTokenizer(tt.s)
		tokens, err := tokenizer.Tokenize()
		if err != nil {
			t.Errorf("Tokenize(%q) = %v, want %v", tt.s, err, nil)
		}
		if got, want := tokens, tt.tokens; !equalTokens(got, want) {
			t.Errorf("Tokenize(%q) = %v, want %v", tt.s, got, want)
		}
	}
}
