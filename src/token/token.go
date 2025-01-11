package token

// Type is a type alias for a string
type Type string

// Token is a value describing the type of token, and its value
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL signifies an unknown token
	ILLEGAL = "ILLEGAL"
	// EOF signifies the end of a file, indicating the parser can
	// stop parsing
	EOF = "EOF"

	// IDENT describes an identifier for a variable
	IDENT = "IDENT"
	// INT describes values that are integers
	INT = "INT"

	// ASSIGN is the operator for assigning values
	ASSIGN = "="
	// PLUS is the operator for addition
	PLUS = "+"

	// COMMA represents a literal comma
	COMMA = ","
	// SEMICOLON represents a literal comma
	SEMICOLON = ";"

	// LPAREN represents a literal comma
	LPAREN = "("
	// RPAREN represents a literal comma
	RPAREN = ")"
	// LBRACE represents a literal comma
	LBRACE = "{"
	// RBRACE represents a literal comma
	RBRACE = "}"

	// FUNCTION is a keyword for declaring functions
	FUNCTION = "FUNCTION"
	// LET is a keyword for declaring variables
	LET = "LET"
)
