package ast

import (
	"fmt"
	// "golang.org/x/text/cases"
)

type TokenizeErrorKind int

const (
	TokenizeErrorKindBraceNotClosed TokenizeErrorKind = iota
	TokenizeErrorKindBracketNotClosed
	TokenizeErrorKindInvalidEscape
)

type TokenizeError struct {
	Kind TokenizeErrorKind
	Pos  int
}

func (t *TokenizeError) Error() string {
	var message = fmt.Sprintf("tokenize error at %d: ", t.Pos)
	switch t.Kind {
	case TokenizeErrorKindBraceNotClosed:
		return message + "brace not closed"
	case TokenizeErrorKindBracketNotClosed:
		return message + "bracket not closed"
	default:
		return message + "unknown error"
	}
}

type Tokenizer struct {
	input  string
	tokens []Token
	pos    int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: input}
}

func (t *Tokenizer) next() string {
	if t.pos >= len(t.input) {
		return ""
	}
	s := string(t.input[t.pos])
	t.pos++
	return s
}

func (t *Tokenizer) Tokenize() ([]Token, error) {
	var s string
	for {
		s = t.next()
		switch s {
		case ".":
			t.tokens = append(t.tokens, Token{Kind: TokenKindDot, Value: "."})
		case "*":
			t.tokens = append(t.tokens, Token{Kind: TokenKindStar, Value: "*"})
		case "+":
			t.tokens = append(t.tokens, Token{Kind: TokenKindPlus, Value: "+"})
		case "?":
			t.tokens = append(t.tokens, Token{Kind: TokenKindQuestion, Value: "?"})
		case "|":
			t.tokens = append(t.tokens, Token{Kind: TokenKindPipe, Value: "|"})
		case "\\":
			token, err := t.tokenizeEscape()
			if err != nil {
				return nil, err
			}
			t.tokens = append(t.tokens, token)
		case "^":
			t.tokens = append(t.tokens, Token{Kind: TokenKindCaret, Value: "^"})
		case "$":
			t.tokens = append(t.tokens, Token{Kind: TokenKindDollar, Value: "$"})
		case "(":
			t.tokens = append(t.tokens, Token{Kind: TokenKindLParen, Value: "("})
		case ")":
			t.tokens = append(t.tokens, Token{Kind: TokenKindRParen, Value: ")"})
		case "{":
			tokens, err := t.tokenizeBrace()
			if err != nil {
				return nil, err
			}
			t.tokens = append(t.tokens, tokens...)
		case "}":
			t.tokens = append(t.tokens, Token{Kind: TokenKindChar, Value: "}"})
		case "[":
			tokens, err := t.tokenizeBracket()
			if err != nil {
				return nil, err
			}
			t.tokens = append(t.tokens, tokens...)
		case "]":
			t.tokens = append(t.tokens, Token{Kind: TokenKindRBracket, Value: "]"})
		case "":
			t.tokens = append(t.tokens, Token{Kind: TokenKindEOF})
			return t.tokens, nil
		default:
			t.tokens = append(t.tokens, Token{Kind: TokenKindChar, Value: s})
		}
	}
}

func (t *Tokenizer) tokenizeEscape() (Token, error) {
	s := t.next()
	if s == "" {
		return Token{}, &TokenizeError{Kind: TokenizeErrorKindInvalidEscape, Pos: t.pos}
	}
	return Token{Kind: TokenKindEscape, Value: "\\" + s}, nil
}

func (t *Tokenizer) tokenizeBrace() ([]Token, error) {
	var tokens []Token
	tokens = append(tokens, Token{Kind: TokenKindLBrace, Value: "{"})
	var number string
	var s string
	for {
		s = t.next()
		switch s {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			number += s
		case ",":
			if number != "" {
				tokens = append(tokens, Token{Kind: TokenKindNumber, Value: number})
			}
			tokens = append(tokens, Token{Kind: TokenKindComma, Value: ","})
			number = ""
		case "}":
			if number != "" {
				tokens = append(tokens, Token{Kind: TokenKindNumber, Value: number})
			}
			tokens = append(tokens, Token{Kind: TokenKindRBrace, Value: "}"})
			return tokens, nil
		default:
			return nil, &TokenizeError{Kind: TokenizeErrorKindBraceNotClosed, Pos: t.pos}
		}
	}
}

func (t *Tokenizer) tokenizeBracket() ([]Token, error) {
	var tokens []Token
	tokens = append(tokens, Token{Kind: TokenKindLBracket, Value: "["})
	var s string
	for {
		s = t.next()
		switch s {
		case "]":
			tokens = append(tokens, Token{Kind: TokenKindRBracket, Value: "]"})
			return tokens, nil
		case "\\":
			token, err := t.tokenizeEscape()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case "-":
			tokens = append(tokens, Token{Kind: TokenKindHyphen, Value: "-"})
		case "^":
			tokens = append(tokens, Token{Kind: TokenKindCaret, Value: "^"})
		default:
			tokens = append(tokens, Token{Kind: TokenKindChar, Value: s})
		}
	}
}
