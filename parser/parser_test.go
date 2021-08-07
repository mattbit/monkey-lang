package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `let a = 1;
let b = a;
let a_snake_case_named_variable = 0;
`
	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	numStatements := len(prog.Statements)
	if numStatements != 3 {
		t.Fatalf("program.Statements should contain 3 statements, got %d",
			numStatements)
	}

	tests := []struct{ expectedIdentifier string }{
		{"a"}, {"b"}, {"a_snake_case_named_variable"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		if !checkLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt is not '*ast.LetStatement', got %T", stmt)
		return false
	}

	if letStmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral is not '%s', got '%s'",
			token.LET, letStmt.TokenLiteral())
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("stmt.Name.Value is not '%s', got '%s'",
			letStmt.Name.Value, name)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
