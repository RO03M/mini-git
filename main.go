package main

import (
	"fmt"
	"mgit/cmd/commands"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("MGit is a custom git command made by Romera :)")
		commands.PrintHelp()
		return
	}

	switch args[0] {
	case "init":
		commands.Init()
	case "add":
		commands.Add(args[1:]...)
	case "commit":
		commands.Commit(args[1])
	case "cat-file":
		commands.CatFile(args[1])
	case "log":
		commands.Log()
	case "checkout":
		commands.Checkout(args[1])
	case "trackable":
		commands.Trackable()
	case "tracked":
		commands.Tracked()
	case "untracked":
		commands.Untracked()
	case "rf":
		commands.RemovedFiles()
	case "help":
		commands.PrintHelp()
	default:
		fmt.Println("Invalid command")
	}
}
