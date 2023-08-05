package main

import (
	"fmt"
	"gnew/repl"
	"os"
)

func main() {
	fmt.Print("This is the Gnew programming language!\nType :q<CR> or exit to stop the repl.\n")
	repl.Start(os.Stdin, os.Stdout)
}
