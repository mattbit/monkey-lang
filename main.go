package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

// BANNER is the Monkey opening banner.
const BANNER = `   __  ___          __
  /  |/  /__  ___  / /_____ __ __
 / /|_/ / _ \/ _ \/  '_/ -_) // /
/_/  /_/\___/_//_/_/\_\\__/\_, /
                          /___/

`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(BANNER)
	fmt.Printf("Hello %s! Welcome to the Monkey REPL.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
