package ast

import "github.com/tkak/go-monkey/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (ls *Identifier) expressionNode()      {}
func (ls *Identifier) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // 'return' token
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }
