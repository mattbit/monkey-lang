package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let forty = 40;
	let two = 2;

	let add = fn(x, y) {
		x + y
	}

	let result = add(forty, two);
	
	!-/*313
	5 < 10 > 5!;

	if (42 > -42) {
		return true;
	} else {
		return false;
	}

	1 == 1;
	1 != 0;
	1 >= 1;
	1 <= 1;
	?=`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// let forty = 40;
		{token.LET, "let"},
		{token.IDENT, "forty"},
		{token.ASSIGN, "="},
		{token.INT, "40"},
		{token.SEMICOLON, ";"},

		// let two = 2;
		{token.LET, "let"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		// let add = fn(x, y) {
		// 	x + y
		// }
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},

		// let result = add(forty, two);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "forty"},
		{token.COMMA, ","},
		{token.IDENT, "two"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// !-/*313
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "313"},

		// 5 < 10 > 5!;
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.BANG, "!"},
		{token.SEMICOLON, ";"},

		// if (42 > -42) {
		// 	return true;
		// } else {
		// 	return false;
		// }
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "42"},
		{token.GT, ">"},
		{token.MINUS, "-"},
		{token.INT, "42"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// 1 == 1;
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},

		// 1 != 0;
		{token.INT, "1"},
		{token.NOT_EQ, "!="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},

		// 1 >= 1;
		{token.INT, "1"},
		{token.GTE, ">="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},

		// 1 <= 1;
		{token.INT, "1"},
		{token.LTE, "<="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},

		// ?=
		{token.ILLEGAL, "?"},
		{token.ASSIGN, "="},

		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
