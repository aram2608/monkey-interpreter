package lexer

import "monkey-interpreter/token"

// defining the lexer struct
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

// Function used to increment our string
func New(input string) *Lexer {
	// Im assuming we initialize a Lexer given the input
	l := &Lexer{input: input}
	// We can then read the characters I presume
	l.readChar()
	// We then return the Lexer
	return l
}

// Helper method used to readh the current character and advance the position
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

// brings in NextToken from the token.go package we made
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

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
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// We handle unexpected characters in our default case
		tok = newToken(token.ILLEGAL, l.ch)
	}

	// We advance forward in our string after every match
	l.readChar()
	return tok
}

// Helper method to return the proper token type
func newToken(tokenType token.TokenType, ch byte) token.Token {
	// We initialize the token given its type and literal value
	return token.Token{Type: tokenType, Literal: string(ch)}
}
