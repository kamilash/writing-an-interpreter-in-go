package main

import (
	"fmt"
	"mycode/repl"
	"os"
)

func main() {
	fmt.Println("Starting REPL of Monkey Language...")
	repl.Start(os.Stdin, os.Stdout)
}
