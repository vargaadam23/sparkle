package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vargaadam23/sparkle/lexer"
	"github.com/vargaadam23/sparkle/token"
)

func Start(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print(">>")
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		lexer := lexer.New(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
