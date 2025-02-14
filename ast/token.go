package ast

type TokenKind int

const (
	TokenKindChar TokenKind = iota
	TokenKindDot
	TokenKindStar
	TokenKindPlus
	TokenKindQuestion
	TokenKindPipe
	TokenKindEscape
	TokenKindCaret
	TokenKindDollar

	TokenKindLParen
	TokenKindRParen
	TokenKindLBrace
	TokenKindRBrace
	TokenKindLBracket
	TokenKindRBracket
	TokenKindEOF

	// Token interpreted in special context
	TokenKindNumber
	TokenKindComma
	TokenKindHyphen
)

type Token struct {
	Kind  TokenKind
	Value string
}
