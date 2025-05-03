package lexer // defines our lexer package

// go test ./... runs all tests in project
// go mod tidy cleans up dependencies
// := is a funky little guy
// it can declare and initilaize a variable in a function in one step

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;` // takes literal strings or something, kinda like lark gramar uses """" grammar """"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input) // you need to explicitly use package.function when using a method

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
