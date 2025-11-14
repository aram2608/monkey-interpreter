package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

const Prompt string = ">> "

// Function to act as the entry point to the REPL
// Reads new lines from the terminal and executes code
// Currently only supports tokenization
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, Prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// We print out the Token struct's fields
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
