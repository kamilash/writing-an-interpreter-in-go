package main

import (
		"mycode/repl"
		"fmt"
		"os"
)

func main() {
		fmt.Println("Starting REPL of Monkey Language...")
		repl.Start(os.Stdin, os.Stdout)
}
