package ast

import (
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ASTKind
type ASTKind int

const (
	ASTKindChar ASTKind = iota
	ASTKindUnion
	ASTKindConcat
	ASTKindStar
	ASTKindGroup
)

type AST struct {
	Kind     ASTKind
	Children []AST
	Value    string
}

func (a AST) String() string {
	switch a.Kind {
	case ASTKindChar:
		return a.Value
	case ASTKindUnion:
		var s string
		for i, child := range a.Children {
			if i > 0 {
				s += "|"
			}
			s += child.String()
		}
		return s
	case ASTKindConcat:
		var s string
		for _, child := range a.Children {
			s += child.String()
		}
		return s
	case ASTKindStar:
		return a.Children[0].String() + "*"
	case ASTKindGroup:
		return "(" + a.Children[0].String() + ")"
	}
	panic("unreachable")
}

func PrintSampleAST() {
	// a(b|c)*d
	ast := AST{
		Kind: ASTKindConcat,
		Children: []AST{
			{
				Kind:  ASTKindChar,
				Value: "a",
			},
			{
				Kind: ASTKindStar,
				Children: []AST{
					{
						Kind: ASTKindGroup,
						Children: []AST{
							{
								Kind: ASTKindUnion,
								Children: []AST{
									{
										Kind:  ASTKindChar,
										Value: "b",
									},
									{
										Kind:  ASTKindChar,
										Value: "c",
									},
								},
							},
						},
					},
					{
						Kind:  ASTKindChar,
						Value: "d",
					},
				},
			},
		},
	}
	fmt.Println(ast.String())
}
