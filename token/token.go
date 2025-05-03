package token // our token package, in go things are defined as packages

// TokenType is defined as a string
// allows us to use many different values as TokenType
// big caveat is that in exchange for flexibility and performance its easier to use
// easy to debug and what not
type TokenType string // string is built in to go itself

// struct defines data types
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y. whatever you feel like
	INT   = "INT"   // numbas

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MODULO   = "%"
	MULTIPLY = "*"
	DIVIDE   = "/"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "()"
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
