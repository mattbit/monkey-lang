package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// A Parser for the Monkey Language.
type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

// New creates a new instance of a Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Initialize the parser by reading the first two token.
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances the parser to the next token.
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses a program and returns the AST.
func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}

		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if p.peekToken.Type != token.IDENT {
		return nil
	}

	p.nextToken()
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if p.peekToken.Type != token.ASSIGN {
		return nil
	}

	// TODO: Skip to the end of the statement.
	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}
