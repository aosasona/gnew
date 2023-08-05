package repl

import (
	"bufio"
	"fmt"
	"gnew/lexer"
	"gnew/token"
	"io"
	"strings"
)

const PROMPT = "\033[36m>>\033[0m "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		if ok := scanner.Scan(); !ok {
			return
		}

		line := scanner.Text()

		if strings.TrimSpace(line) == ":q" || strings.TrimSpace(line) == "exit" {
			fmt.Fprintf(out, "Goodbye!\n")
			return
		}

		l := lexer.New(line)

		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tk)
		}
	}
}
