package main

import (
	"fmt"
	"os"

	command "github.com/santoshbachar/navik-cli/command"
)

func main() {
	fmt.Println("Hello... starting Navik CLI")

	args := os.Args

	fmt.Println("all args", args)

	command.Serve(args[1:])

}
