package lexer

import (
	"monkey/token"
)

func New(input string) *lexer {
	return &lexer{input: input}
}

type lexer struct {
	input        string
	chPosition   int
	readPosition int
	ch           byte
}

func (l *lexer) NextToken() token.Token {
	var literal string
	var tokenType token.Type

	if l.chPosition >= len(l.input) {
		literal = ""
	} else {
		l.ch = l.input[l.chPosition]
		literal = string(l.ch)
	}

	switch literal {
	case "=":
		tokenType = token.ASSIGN
	case "+":
		tokenType = token.PLUS
	case "(":
		tokenType = token.LPAREN
	case ")":
		tokenType = token.RPAREN
	case "{":
		tokenType = token.LBRACE
	case "}":
		tokenType = token.RBRACE
	case ",":
		tokenType = token.COMMA
	case ";":
		tokenType = token.SEMICOLON
	case "":
		tokenType = token.EOF
	}

	result := token.Token{Type: tokenType, Literal: literal}
	l.chPosition += 1

	return result
}
