package main

import (
	"fmt"
	"mgit/cmd/commands"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Custom git command made by Romera :)")
		fmt.Println("Run help to see the available commands")
		return
	}

	switch args[0] {
	case "init":
		commands.Init()
	case "add":
		commands.Add(args[1])
	case "commit":
		commands.Commit(args[1])
	case "cat-file":
		commands.CatFile(args[1])
	case "log":
		commands.Log()
	case "checkout":
		commands.Checkout(args[1])
	default:
		fmt.Println("Invalid command")
	}
}
