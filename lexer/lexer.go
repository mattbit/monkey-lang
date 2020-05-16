package lexer

import "monkey/token"

// Lexer is the Monkey lang lexer.
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New returns a Lexer for the given input string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken reads the next token and advances the lexer.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type, tok.Literal = token.ASSIGN, string(l.ch)
		}
	case '+':
		tok.Type, tok.Literal = token.PLUS, string(l.ch)
	case '-':
		tok.Type, tok.Literal = token.MINUS, string(l.ch)
	case '*':
		tok.Type, tok.Literal = token.ASTERISK, string(l.ch)
	case '/':
		tok.Type, tok.Literal = token.SLASH, string(l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type, tok.Literal = token.BANG, string(l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.LTE
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type, tok.Literal = token.LT, string(l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.GTE
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type, tok.Literal = token.GT, string(l.ch)
		}
	case ',':
		tok.Type, tok.Literal = token.COMMA, string(l.ch)
	case ';':
		tok.Type, tok.Literal = token.SEMICOLON, string(l.ch)
	case '(':
		tok.Type, tok.Literal = token.LPAREN, string(l.ch)
	case ')':
		tok.Type, tok.Literal = token.RPAREN, string(l.ch)
	case '{':
		tok.Type, tok.Literal = token.LBRACE, string(l.ch)
	case '}':
		tok.Type, tok.Literal = token.RBRACE, string(l.ch)
	case 0:
		tok.Type, tok.Literal = token.EOF, ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}

		tok.Type, tok.Literal = token.ILLEGAL, string(l.ch)
	}

	l.readChar()
	return tok
}
