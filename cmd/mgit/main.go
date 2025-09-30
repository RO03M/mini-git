package main

import (
	"fmt"
	"mgit/internal/commands"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("MGit is a custom git command made by Romera :)")
		return
	}

	command := args[0]

	switch command {
	case "init":
		commands.Init()
	case "add":
		files := args[1:]
		commands.Add(files...)
	case "rm":
		files := args[1:]
		commands.Rm(files...)
	case "log":
		commands.Log()
	case "cat-file", "cat":
		commands.Catfile(args[1])
	case "status":
		commands.Status()
	case "commit":
		commands.Commit(args[1:]...)
	case "trackable":
		commands.Trackable()
	case "rev-parse", "revparse":
		commands.RevParse(args[1:]...)
	case "branch":
		commands.Branch(args[1:]...)
	case "switch":
		commands.Switch(args[1:]...)
	case "checkout":
		commands.Checkout(args[1:]...)
	case "help":
		commands.Help()
	case "version", "-v", "--version", "v":
		commands.Version()
	case "diff":
		commands.Diff(args[1:]...)
	default:
		fmt.Println("Invalid command")
	}
}
