package ast

import "monkey/token"

// Node is the interface for all AST nodes.
type Node interface {
	TokenLiteral() string
}

// Statement is a Node that does not return a value.
type Statement interface {
	Node
	statementNode()
}

// Expression is a Node that returns a value.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the token for debugging.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// An Identifier is the name of a variable.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expresssionNode() {}

// TokenLiteral returns the literal value of the token for debugging
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement represents the assignment of a variable.
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns the literal value of the token for debugging
func (s *LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *LetStatement) statementNode() {

}
