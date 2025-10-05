package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mtzack9/the-interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Println(repl.INTERPRETER_FACE)
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
