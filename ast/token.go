package ast

//go:generate go run github.com/dmarkham/enumer -type=TokenKind
type TokenKind int

const (
	TokenKindChar TokenKind = iota
	TokenKindUnion
	TokenKindStar
	TokenKindLParen
	TokenKindRParen
	TokenKindEOF
)
