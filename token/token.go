package token // our token package, in go things are defined as packages

// TokenType is defined as a string
type TokenType string

// We create a Token struct that stores the meta data associated with each token
type Token struct {
	Type    TokenType
	Literal string
}

// We define a set of constant token Type
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	// identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

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

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
