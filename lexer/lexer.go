package lexer

import "monkey-interpreter/token"

// defining the lexer structure
type Lexer struct {
	input        string
	position     int  // current position in input (points to current character)
	readPosition int  // cuurent reading position in input (after the current character)
	ch           byte // current character under examination
}

// some sort of function that i do not comprehend
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // reads characters and increments, the func we made below!
	return l
}

// a helper function to give us the next character and advance our position
// checks whether we have reach the end of input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition] // accesses the next position
	}
	l.position = l.readPosition
	l.readPosition += 1 // increments along the input by 1
}

// brings in NextToken from the token.go package we made
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
		// handles unexpected characters
		tok = newToken(token.ILLEGAL, l.ch)
	}
	return tok
}

// method to use newToken
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
