package main

import (
	"fmt"
	"os"
	"os/user"
	"polieren/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! welcome to polieren repl\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
