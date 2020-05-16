package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey REPL.\n", user.Username)
	fmt.Printf("Type some Mokey code.\n")
	repl.Start(os.Stdin, os.Stdout)
}
