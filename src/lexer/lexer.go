package lexer

import (
	"monkey/token"
)

func New(input string) *lexer {
	l := &lexer{input: input}
	l.readChar()

	return l
}

type lexer struct {
	input        string
	chPosition   int
	readPosition int
	ch           byte
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // represents NUL in ASCII
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.chPosition = l.readPosition // chPosition always points to the current char position
	l.readPosition += 1           // readPosition always points to next char to read
}

func (l *lexer) NextToken() token.Token {
	var tokenType token.Type
	literal := string(l.ch)

	switch l.ch {
	case '=':
		tokenType = token.ASSIGN
	case '+':
		tokenType = token.PLUS
	case '(':
		tokenType = token.LPAREN
	case ')':
		tokenType = token.RPAREN
	case '{':
		tokenType = token.LBRACE
	case '}':
		tokenType = token.RBRACE
	case ',':
		tokenType = token.COMMA
	case ';':
		tokenType = token.SEMICOLON
	case 0:
		tokenType = token.EOF
		literal = ""
	}

	result := token.Token{Type: tokenType, Literal: literal}

	l.readChar()

	return result
}
