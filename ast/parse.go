package ast

import (
	"errors"
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

type ParseErrorKind int

const (
	ParseErrorKindParenNotClosed ParseErrorKind = iota
	ParseErrorKindParenNotOpened
	ParseErrorKindUnexpectedPipe
	ParseErrorKindUnexpectedStar
	ParseErrorKindUnexpectedEOF

	ParseErrorKindInvalidAST
)

type ParseError struct {
	Kind ParseErrorKind
	Pos  int
}

func (p *Parser) newParseError(kind ParseErrorKind) *ParseError {
	return &ParseError{Kind: kind, Pos: p.pos}
}

func (p *ParseError) Error() string {
	// エラーの種類に関するメッセージと、エラーが発生した位置を返す
	var message = fmt.Sprintf("parse error at %d: ", p.Pos)
	switch p.Kind {
	case ParseErrorKindParenNotClosed:
		return message + "parenthesis not closed"
	case ParseErrorKindParenNotOpened:
		return message + "parenthesis not opened"
	case ParseErrorKindUnexpectedPipe:
		return message + "unexpected pipe"
	case ParseErrorKindUnexpectedStar:
		return message + "unexpected star"
	case ParseErrorKindUnexpectedEOF:
		return message + "unexpected EOF"
	case ParseErrorKindInvalidAST:
		return message + "invalid AST"
	}
	return message + "unknown error"
}

func (err *ParseError) CheckKind(kind ParseErrorKind) bool {
	return err.Kind == kind
}

func (p *Parser) peek() Token {
	return p.tokens[p.pos]
}

func (p *Parser) next() Token {
	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) expect(kind TokenKind) error {
	token := p.next()
	switch token.Kind {
	case kind:
		return nil
	case TokenKindEOF:
		return p.newParseError(ParseErrorKindUnexpectedEOF)
	case TokenKindRParen:
		return p.newParseError(ParseErrorKindParenNotOpened)
	case TokenKindStar:
		return p.newParseError(ParseErrorKindUnexpectedStar)
	default:
		return p.newParseError(ParseErrorKindInvalidAST)
	}
}

func (p *Parser) Parse() (*AST, error) {
	ast, err := p.parseUnion()
	if err != nil {
		return nil, err
	}
	token := p.peek()
	switch token.Kind {
	case TokenKindEOF:
		return ast, nil
	case TokenKindRParen:
		return nil, p.newParseError(ParseErrorKindParenNotOpened)
	case TokenKindStar:
		return nil, p.newParseError(ParseErrorKindUnexpectedStar)
	case TokenKindPipe:
		return nil, p.newParseError(ParseErrorKindUnexpectedPipe)
	default:
		return nil, p.newParseError(ParseErrorKindInvalidAST)
	}
}

func (p *Parser) parseUnion() (*AST, error) {
	ast := NewUnionAST(nil)
	for {
		concat, err := p.parseConcat()
		if err != nil {
			return nil, err
		}
		ast.Children = append(ast.Children, concat)
		token := p.peek()
		if token.Kind != TokenKindPipe {
			break
		}
		p.next()
	}
	if len(ast.Children) == 0 {
		return nil, p.newParseError(ParseErrorKindInvalidAST)
	} else if len(ast.Children) == 1 {
		return ast.Children[0], nil
	}
	return ast, nil
}

func (p *Parser) parseConcat() (*AST, error) {
	children := []*AST{}
	var ast *AST
	var err error
	for {
		ast, err = p.parsePostfix()
		if err != nil {
			break
		}
		children = append(children, ast)
	}
	if len(children) == 0 {
		return nil, err
	} else if len(children) == 1 {
		return children[0], nil
	} else {
		return NewConcatAST(children), nil
	}
}

func (p *Parser) parsePostfix() (*AST, error) {
	ast, err := p.parseGroup()
	token := p.peek()
	if err != nil {
		var parseError *ParseError
		if errors.As(err, &parseError) {
			switch token.Kind {
			case TokenKindStar:
				return nil, p.newParseError(ParseErrorKindUnexpectedStar)
			case TokenKindPipe:
				return nil, p.newParseError(ParseErrorKindUnexpectedPipe)
			case TokenKindEOF:
				return nil, p.newParseError(ParseErrorKindUnexpectedEOF)
			case TokenKindRParen:
				return nil, p.newParseError(ParseErrorKindParenNotOpened)
			case TokenKindLParen:
				return nil, p.newParseError(ParseErrorKindInvalidAST)
			case TokenKindChar:
				ast = NewCharAST(token.Value)
				p.next()
			}
		} else {
			return nil, err
		}
	}
	token = p.peek()
	if token.Kind == TokenKindStar {
		p.next()
		return NewStarAST(ast), nil
	}
	return ast, nil
}

func (p *Parser) parseGroup() (*AST, error) {
	token := p.peek()
	if token.Kind != TokenKindLParen {
		return nil, p.newParseError(ParseErrorKindParenNotOpened)
	}
	p.next()
	ast, err := p.parseUnion()
	if err != nil {
		return nil, err
	}
	token = p.peek()
	switch token.Kind {
	case TokenKindRParen:
		p.next()
		return NewGroupAST(ast), nil
	case TokenKindEOF:
		return nil, p.newParseError(ParseErrorKindParenNotClosed)
	case TokenKindStar:
		return nil, p.newParseError(ParseErrorKindUnexpectedStar)
	case TokenKindPipe:
		return nil, p.newParseError(ParseErrorKindUnexpectedPipe)
	default:
		return nil, p.newParseError(ParseErrorKindInvalidAST)
	}
}
