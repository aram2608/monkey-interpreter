package lexer

import (
	"monkey-interpreter/token"
)

// The Lexer object for Moneky
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// Function used to create the lexer, initilaizing it with the input
// returning a pointer to the Lexer
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// Helper method used to read the current character and advance the position
func (l *Lexer) readChar() {
	// We need to make sure the read position is not out of bounds
	if l.isEnd() {
		// If it is we set the current char to 0 to catch the end of file
		l.ch = 0
	} else {
		// Otherwise we advance foward
		l.ch = l.input[l.readPosition]
	}
	// We swap the current position and increment forward
	l.position = l.readPosition
	l.readPosition += 1
}

// Helper method to peek at the current character, returns a byte
func (l *Lexer) peekChar() byte {
	// We catch out of bounds indexing
	if l.isEnd() {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Helper method to test if we are at the end of the file, returns a boolean
func (l *Lexer) isEnd() bool {
	return l.readPosition >= len(l.input)
}

// Helper method used to match each character to a token type
func (l *Lexer) NextToken() token.Token {
	// We forward declare a token
	var tok token.Token

	l.skipWhiteSpace()

	// We catch the new character and match it to each token type case
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NotEQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
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
