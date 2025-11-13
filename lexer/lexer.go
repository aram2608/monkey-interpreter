package lexer

import (
	"monkey-interpreter/token"
)

// We define the Lexer struct
type Lexer struct {
	// The source code/string
	input string
	// The current position
	position int
	// The current position we are reading
	readPosition int
	// The current character as a byte
	ch byte
}

// Function used to create the lexer, initilaizing it with the input
// returning a pointer to the Lexer
func New(input string) *Lexer {
	// We initialize a Lexer struct in place and return a reference, assigning it
	// to l
	l := &Lexer{input: input}
	// We can then read the character
	l.readChar()
	// We then return the Lexer
	return l
}

// Helper method used to read the current character and advance the position
func (l *Lexer) readChar() {
	// We need to make sure the read position is not out of bounds
	if l.readPosition >= len(l.input) {
		// If it is we set the current char to 0 to catch the end of file
		l.ch = 0
	} else {
		// Otherwise we advance foward
		l.ch = l.input[l.readPosition]
	}
	// We swap the current position
	l.position = l.readPosition
	// We can then increment the read position forward
	l.readPosition += 1
}

// Helper method used to match each character to a token type
func (l *Lexer) NextToken() token.Token {
	// We forward declare a token
	var tok token.Token

	l.skipWhiteSpace()

	// We catch the new character and match it to each token type case
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '%':
		tok = newToken(token.MODULO, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// An EOF is denoted by an 0 in our Lexer
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// We catch identifiers and keywords
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			// We handle unexpected characters in our default case
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// We advance forward in our string after every match
	l.readChar()
	return tok
}

// Helper function to skip past white space
// Whitespace is meaningless in Monkey and is used just for formatting
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

// Function to read identifiers
// returns the substring containing the identifier's lexeme
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Helper method to reat identifiers
// returns the substring containing the number's lexeme
func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Helper function to test if a current character is a letter or an underscore
func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

// Helper function to test if a current character is a number
func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// Helper function used to create a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	// We initialize the token given its type and literal value
	return token.Token{Type: tokenType, Literal: string(ch)}
}
