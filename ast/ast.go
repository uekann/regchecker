package ast

// "fmt"

type ASTKind int

//go:generate enumer -type=ASTKind -json
const (
	ASTKindChar ASTKind = iota
	ASTKindUnion
	ASTKindConcat
	ASTKindStar
	ASTKindGroup
)

type AST struct {
	Kind     ASTKind
	Children []*AST
	Value    string
}

func (a *AST) String() string {
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

// func (a *AST) DebugString() string {
// 	var children []string
// 	for _, child := range a.Children {
// 		children = append(children, child.DebugString())
// 	}

func NewCharAST(value string) *AST {
	return &AST{Kind: ASTKindChar, Value: value}
}

func NewUnionAST(children []*AST) *AST {
	return &AST{Kind: ASTKindUnion, Children: children}
}

func NewConcatAST(children []*AST) *AST {
	return &AST{Kind: ASTKindConcat, Children: children}
}

func NewStarAST(child *AST) *AST {
	return &AST{Kind: ASTKindStar, Children: []*AST{child}}
}

func NewGroupAST(child *AST) *AST {
	return &AST{Kind: ASTKindGroup, Children: []*AST{child}}
}
