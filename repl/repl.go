package repl

import (
	"bufio"
	"fmt"
	"io"
	"mlang/lexer"
	"mlang/token"
)

const PROMPT = "$> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Fprintf(output, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for tok := lexer.Next(); tok.TYPE != token.EOF; tok = lexer.Next() {
			fmt.Fprintf(output, "%+v\n", tok)
		}

	}
}
